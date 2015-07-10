package rpcapi

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"testing"
	"thrift/gen-go/rpc"
	"time"
)

const TIMEOUT = time.Second * 15

func Test_CallEchoService(t *testing.T) {
	socket, err := thrift.NewTSocketTimeout(net.JoinHostPort("127.0.0.1", "8080"), TIMEOUT)
	if err != nil {
		t.Fatal("Unable to new client socket", err)
	}

	transport := thrift.NewTFramedTransport(socket)
	var protocol thrift.TProtocol = thrift.NewTBinaryProtocolTransport(transport)
	protocol = thrift.NewTMultiplexedProtocol(protocol, "EchoService")
	client := rpc.NewEchoServiceClientProtocol(transport, protocol, protocol)

	err = transport.Open()
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

func Benchmark_CallEchoService(b *testing.B) {
	socket, err := thrift.NewTSocketTimeout(net.JoinHostPort("127.0.0.1", "8080"), TIMEOUT)
	if err != nil {
		t.Fatal("Unable to new client socket", err)
	}

	transport := thrift.NewTFramedTransport(socket)
	var protocol thrift.TProtocol = thrift.NewTBinaryProtocolTransport(transport)
	protocol = thrift.NewTMultiplexedProtocol(protocol, "EchoService")
	client := rpc.NewEchoServiceClientProtocol(transport, protocol, protocol)

	err = transport.Open()
	if err != nil {
		t.Fatal("Unable to open client socket", err)
	}
	defer transport.Close()

	for i := 0; i < b.N; i++ { //use b.N for looping
		_, err2 := client.Echo(currentTimeMillis(), "hello world")
		if err2 != nil {
			exception := thrift.NewTTransportExceptionFromError(err2)
			t.Fatal("test TestCallEchoService failed", exception.TypeId(), exception.Err())
		}
	}
}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
