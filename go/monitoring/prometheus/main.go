package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
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

	go func() {
		for i := 0; i < 100; i++ {
			reqID := uuid.New().String()
			start := time.Now()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			duration := time.Since(start)
			fmt.Printf("request %s took %dms\n", reqID, int64(duration/time.Millisecond))
			labels := prometheus.Labels{"method": "GET", "path": "/foo"}
			// we cast the histogram into a ExemplarObserver in order to call ObserveWithExemplar on it
			ReqHistogram.With(labels).(prometheus.ExemplarObserver).ObserveWithExemplar(duration.Seconds(), prometheus.Labels{"request_id": reqID})
		}
	}()
	s.ListenAndServe()
}

func addMonitor(mux *http.ServeMux) error {
	mux.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{EnableOpenMetrics: true}))

	return nil
}
