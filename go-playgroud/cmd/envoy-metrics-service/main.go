package main

import (
	"log"
	"net"
	"os"

	metricsservice "github.com/envoyproxy/go-control-plane/envoy/service/metrics/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ metricsservice.MetricsServiceServer = (*server)(nil)

type server struct {
}

func (s *server) StreamMetrics(mtserver metricsservice.MetricsService_StreamMetricsServer) error {
	log.Println("Inside metrics service log")
	messge, err := mtserver.Recv()
	if err != nil {
		log.Fatalf("failed to receive a message: %v", err)
		return err
	}

	envoyMetrics := messge.GetEnvoyMetrics()
	envoyIdentifier := messge.GetIdentifier().String()
	log.Printf("Received metrics from %s", envoyIdentifier)
	for _, metric := range envoyMetrics {
		log.Printf("Received metric: %v", metric)
	}

	return nil
}

func main() {
	address := os.Getenv("LISTEN_ADDRESS")
	if address == "" {
		address = "0.0.0.0:8080"
	}

	// create a TCP listener on port 8080
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr())

	s := grpc.NewServer()
	metricsservice.RegisterMetricsServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
