package main

import (
	"fmt"
	"net/http"
	"time"

	ocprom "contrib.go.opencensus.io/exporter/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"go.opencensus.io/stats/view"
	"istio.io/istio/pkg/monitoring"
)

var (
	serverStart = time.Now()
)

func addMonitor(mux *http.ServeMux) error {
	exporter, err := ocprom.NewExporter(ocprom.Options{Registry: prometheus.DefaultRegisterer.(*prometheus.Registry)})
	if err != nil {
		return fmt.Errorf("could not set up prometheus exporter: %v", err)
	}
	view.RegisterExporter(exporter)
	mux.Handle("/metrics", exporter)

	return nil
}

func main() {
	units := monitoring.CreateLabel("units")
	dg1 := monitoring.NewDerivedGauge(
		"test_derived_gauge",
		"",
	)

	dg1.ValueFrom(
		func() float64 {
			return float64(time.Since(serverStart).Seconds())
		},
		units.Value("second"),
	)

	dg1.ValueFrom(
		func() float64 {
			return float64(time.Since(serverStart).Minutes())
		},
		units.Value("minute"),
	)

	mux := http.NewServeMux()
	if err := addMonitor(mux); err != nil {
		fmt.Printf("could not establish self-monitoring: %v\n", err)
	}

	s := &http.Server{
		Addr:    ":15014",
		Handler: mux,
	}
	fmt.Println("start to listen :15014")
	s.ListenAndServe()
}
