package rpcapi

import (
	"fmt"
)

type EchoServiceImpl struct {
}

func (this *EchoServiceImpl) Echo(callTime int64, helloCode string) (r string, err error) {
	fmt.Println("-->Echo:", callTime, helloCode)
	// r = fmt.Sprintf("%s Gopher", helloCode)
	return
}
