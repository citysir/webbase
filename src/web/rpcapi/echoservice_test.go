package rpcapi

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"io"
	"log"
	"net"
	"testing"
	"thrift/gen-go/rpc"
	"time"
)

func TestCallEchoService(t *testing.T) {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "8080"))
	if err != nil {
		log.Fatalln("error resolving address:", err)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewEchoServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening socket to 127.0.0.1:8080", err)
	}
	defer transport.Close()

	r, err := client.Echo(currentTimeMillis(), "hello world")
	if err != nil && err != io.EOF {
		t.Error("test TestCallEchoService failed", err, err.Error())
	} else {
		t.Log("test TestCallEchoService pass:", r)
	}
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
