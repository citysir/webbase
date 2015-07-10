package app

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"thrift/gen-go/rpc"
	"time"
	"web/rpcapi"
)

const (
	TIMEOUT = time.Hour * 8
)

func startRpcServe(port string) {
	processor := thrift.NewTMultiplexedProcessor()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	serverTransport, err := thrift.NewTServerSocketTimeout(fmt.Sprintf(":%s", port), TIMEOUT)
	if err != nil {
		log.Fatalln("Unable to create server socket", err)
	}

	echoProcessor := rpc.NewEchoServiceProcessor(&rpcapi.EchoServiceImpl{})
	processor.RegisterProcessor("echo", echoProcessor)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	server.Serve()
}
