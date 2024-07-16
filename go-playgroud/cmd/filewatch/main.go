package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"istio.io/istio/pkg/file"
	"istio.io/istio/pkg/filewatcher"
)

var logger = log.Default()

func main() {
	watchedFile := os.Getenv("WATCHED_FILE")
	if watchedFile == "" {
		logger.Println("WATCHED_FILE env var not set")
		os.Exit(1)
	}
	w := filewatcher.NewWatcher()
	defer w.Close()

	if err := w.Add(watchedFile); err != nil {
		logger.Printf("error adding file to watch: %v\n", err)
		os.Exit(1)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	events := w.Events(watchedFile)
	for {

		select {
		case e := <-events:
			logger.Printf("file %s modified: %v\n", watchedFile, e)
			b, _ := os.ReadFile(watchedFile)
			logger.Printf("file %s content: %s\n", watchedFile, string(b))
		case <-sigCh:
			logger.Println("received signal, exiting")
			return
		}
	}
}

type Watcher struct {
	watcher *fsnotify.Watcher
	Events  chan struct{}
	Errors  chan error
}

func CreateFileWatcher(paths ...string) (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("watcher create: %v", err)
	}

	fileModified, errChan := make(chan struct{}), make(chan error)
	go watchFiles(watcher, fileModified, errChan)

	for _, path := range paths {
		if !file.Exists(path) {
			logger.Printf("file watcher skipping watch on non-existent path: %v\n", path)
			continue
		}
		if err := watcher.Add(path); err != nil {
			if closeErr := watcher.Close(); closeErr != nil {
				err = fmt.Errorf("%s: %w", closeErr.Error(), err)
			}
			return nil, err
		}
	}

	return &Watcher{
		watcher: watcher,
		Events:  fileModified,
		Errors:  errChan,
	}, nil
}

func watchFiles(watcher *fsnotify.Watcher, fileModified chan struct{}, errChan chan error) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&(fsnotify.Create|fsnotify.Write|fsnotify.Remove) != 0 {
				logger.Printf("file modified: %v\n", event.Name)
				fileModified <- struct{}{}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			errChan <- err
		}
	}
}
