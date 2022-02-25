package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/zirain/thrift-demo/helloworld/pkg/gen-go/rpc"
)

func handleClient(client *rpc.GreetingServiceClient) error {
	str, err := client.SayHello(context.TODO(), "zirain")
	fmt.Println(str)
	return err
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocket(addr)

	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport, _ = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	return handleClient(rpc.NewGreetingServiceClientFactory(transport, protocolFactory))
}

func main() {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	runClient(transportFactory, protocolFactory, "localhost:18000")
}
