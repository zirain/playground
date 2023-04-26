package main

import (
	"context"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(context.TODO(), "GET", "www.google.com", nil)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	// lint will fail
	// defer resp.Body.Close()
	defer func(b io.ReadCloser) {
		_ = b.Close()
	}(resp.Body)
}
