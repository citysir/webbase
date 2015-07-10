package app

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func Run(webPort, rpcPort string) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	printEnvironments(webPort, rpcPort)

	go startWebServe(webPort)
	go startRpcServe(rpcPort)

	stop()
}

func printEnvironments(webPort, rpcPort string) {
	fmt.Println("os:", runtime.GOOS, runtime.GOARCH)
	fmt.Println("listen:", webPort, rpcPort)
}

func stop() {
	chSignal := make(chan os.Signal)
	signal.Notify(chSignal, os.Kill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSignal)

	fmt.Println("stopped.")
}
