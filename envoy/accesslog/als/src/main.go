package main

import (
	"log"
	"net"
	"net/http"

	alsv2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	alsv3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"google.golang.org/grpc"
)

var (
	LogCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "log_count",
		Help: "The total number of logs received.",
	}, []string{"api_version", "log_type"})
)

func init() {
	// Register the summary and the histogram with Prometheus's default registry.
	prometheus.MustRegister(LogCount)
}

type ALSServer struct {
}

func (a *ALSServer) StreamAccessLogs(logStream alsv2.AccessLogService_StreamAccessLogsServer) error {
	log.Println("Streaming als v2 logs")
	for {
		data, err := logStream.Recv()
		if err != nil {
			return err
		}

		log.Println("Received v2 log data from client:" + data.GetIdentifier().String())

		httpLogs := data.GetHttpLogs()
		if httpLogs != nil {
			LogCount.WithLabelValues("v2", "http").Add(float64(len(httpLogs.LogEntry)))
			for _, logEntry := range httpLogs.LogEntry {
				log.Println("http: " + logEntry.String())
			}
		}

		tcpLogs := data.GetTcpLogs()
		if tcpLogs != nil {
			LogCount.WithLabelValues("v2", "tcp").Add(float64(len(tcpLogs.LogEntry)))
			for _, logEntry := range tcpLogs.LogEntry {
				log.Println("tcp: " + logEntry.String())
			}
		}
	}
}

type ALSServerV3 struct {
}

func (a *ALSServerV3) StreamAccessLogs(logStream alsv3.AccessLogService_StreamAccessLogsServer) error {
	log.Println("Streaming als v3 logs")
	for {
		data, err := logStream.Recv()
		if err != nil {
			return err
		}

		log.Println("Received v3 log data from client:" + data.GetIdentifier().String())

		httpLogs := data.GetHttpLogs()
		if httpLogs != nil {
			LogCount.WithLabelValues("v3", "http").Add(float64(len(httpLogs.LogEntry)))
			for _, logEntry := range httpLogs.LogEntry {
				log.Println("http: " + logEntry.String())
			}
		}

		tcpLogs := data.GetTcpLogs()
		if tcpLogs != nil {
			LogCount.WithLabelValues("v3", "tcp").Add(float64(len(tcpLogs.LogEntry)))
			for _, logEntry := range tcpLogs.LogEntry {
				log.Println("tcp: " + logEntry.String())
			}
		}
	}
}

func NewALSServer() *ALSServer {
	return &ALSServer{}
}

func NewALSServerV3() *ALSServerV3 {
	return &ALSServerV3{}
}

func main() {
	mux := http.NewServeMux()
	if err := addMonitor(mux); err != nil {
		log.Printf("could not establish self-monitoring: %v\n", err)
	}

	s := &http.Server{
		Addr:    ":19001",
		Handler: mux,
	}

	go func() {
		log.Println("Starting prometheus server on port 19001")
		s.ListenAndServe()
	}()

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to start listener on port 8080: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	alsv2.RegisterAccessLogServiceServer(grpcServer, NewALSServer())
	alsv3.RegisterAccessLogServiceServer(grpcServer, NewALSServerV3())
	log.Println("Starting ALS Server")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("grpc serve err: %v", err)
	}
}

func addMonitor(mux *http.ServeMux) error {
	mux.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{EnableOpenMetrics: true}))

	return nil
}