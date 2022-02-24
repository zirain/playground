package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	mux := http.NewServeMux()
	if err := addMonitor(mux); err != nil {
		fmt.Println("could not establish self-monitoring: %v", err)
	}

	s := &http.Server{
		Addr:    ":15014",
		Handler: mux,
	}
	fmt.Println("start to listen :15014")
	s.ListenAndServe()
}

func addMonitor(mux *http.ServeMux) error {
	mux.Handle("/metrics", promhttp.Handler())

	return nil
}
