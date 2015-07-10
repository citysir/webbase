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

func registerProcessors(processor *TMultiplexedProcessor) {
	processor.RegisterProcessor("EchoService", rpc.NewEchoServiceProcessor(&rpcapi.EchoServiceImpl{}))
}

func startRpcServe(port string) {
	socket, err := thrift.NewTServerSocketTimeout(fmt.Sprintf(":%s", port), TIMEOUT)
	if err != nil {
		log.Fatalln("Unable to create server socket", err)
	}

	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	transport := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	processor := thrift.NewTMultiplexedProcessor()

	registerProcessors(processor)

	server := thrift.NewTSimpleServer4(processor, socket, transport, protocol)
	server.Serve()
}
