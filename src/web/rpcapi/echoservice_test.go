package rpcapi

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"testing"
	"thrift/gen-go/rpc"
	"time"
)

const TIMEOUT = time.Second * 15

func TestCallEchoService(t *testing.T) {
	socket := thrift.NewTSocketTimeout(net.JoinHostPort("127.0.0.1", "8080"), TIMEOUT)
	transport := thrift.NewTFramedTransport(socket)
	var protocol thrift.TProtocol = thrift.NewTBinaryProtocolTransport(transport)
	protocol = thrift.NewTMultiplexedProtocol(protocol, "EchoService")
	client := rpc.NewEchoServiceClientProtocol(transport, protocol, protocol)

	err := transport.Open()
	if err != nil {
		t.Fatal("Unable to open client socket", err)
	}
	defer transport.Close()

	r, err := client.Echo(currentTimeMillis(), "hello world")
	if err != nil {
		exception := thrift.NewTTransportExceptionFromError(err)
		t.Error("test TestCallEchoService failed", exception.TypeId(), exception.Err())
	} else {
		t.Log("test TestCallEchoService pass:", r)
	}
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
