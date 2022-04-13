package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	GaugeVecApiDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apiDuration",
		Help: "api耗时单位ms",
	}, []string{"WSorAPI"})
	GaugeVecApiMethod = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apiCount",
		Help: "各种网络请求次数",
	}, []string{"method"})
	GaugeVecApiError = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apiErrorCount",
		Help: "请求api错误的次数type: api/ws",
	}, []string{"type"})

	ReqHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_requests_duration_second",
			Help: "Time to execute http requests",
			Buckets: []float64{
				0.05, 0.1, 0.2, 0.3, 0.4, 0.5,
				0.6, 0.7, 0.8, 0.9},
		},
		[]string{"method", "path"})
)

func init() {
	// Register the summary and the histogram with Prometheus's default registry.
	prometheus.MustRegister(GaugeVecApiMethod, GaugeVecApiDuration, GaugeVecApiError, ReqHistogram)
}
