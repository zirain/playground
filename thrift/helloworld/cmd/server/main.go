package main

import (
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/zirain/thrift-demo/helloworld/pkg/gen-go/rpc"
	"github.com/zirain/thrift-demo/helloworld/pkg/service"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var err error
	transport, err := thrift.NewTServerSocket(addr)

	if err != nil {
		fmt.Println(err)
	}
	handler := service.NewGreetingService()
	processor := rpc.NewGreetingServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}

func main() {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	runServer(transportFactory, protocolFactory, "localhost:9090")
}
