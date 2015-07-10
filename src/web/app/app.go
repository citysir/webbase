package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
	"time"
	"web/util"
)

func Run(webPort, rpcPort string) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	printEnvironments(webPort)

	go startWebServe(webPort)
	go startRpcServe(rpcPort)

	stop()
}

func printEnvironments(port string) {
	fmt.Println("os:", runtime.GOOS, runtime.GOARCH)
	fmt.Println("listen:", port)
}

func stop() {
	chSignal := make(chan os.Signal)
	signal.Notify(chSignal, os.Kill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSignal)

	fmt.Println("stopped.")
}
