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

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	echoProcessor := rpc.NewEchoServiceProcessor(&rpcapi.EchoServiceImpl{})
	processor.RegisterProcessor("EchoService", echoProcessor)

	server.Serve()
}
