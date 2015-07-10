package app

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"time"
	"web/rpcapi"
)

const (
	TIMEOUT = time.Second * 5
)

func startRpcServe(port string) {
	processor := thrift.NewTMultiplexedProcessor()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	serverTransport, err := thrift.NewTServerSocketTimeout(fmt.Sprintf(":%s", port), TIMEOUT)
	if err != nil {
		log.Fatalln("Unable to create server socket", err)
	}

	server = thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	echoProcessor := multiplexedprotocoltest.NewEchoServiceProcessor(&rpcapi.EchoServiceImpl{})
	processor.RegisterProcessor("EchoService", echoProcessor)

	server.Serve()
}
