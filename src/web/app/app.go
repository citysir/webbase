package app

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"syscall"
	"time"
	"web/util"
)

type handler struct {
}

func (this *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var httpStatus int
	if pathFunc, present := pathMap[r.URL.Path]; present {
		httpStatus = safeRun(w, r, pathFunc)
	} else {
		httpStatus = http.StatusNotFound
		http.NotFound(w, r)
	}
}

func Run() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	printEnvironments()

	go startServe()

	stop()
}

func printEnvironments() {
	fmt.Println("os:", runtime.GOOS, runtime.GOARCH)
}

func safeRun(w http.ResponseWriter, r *http.Request, realFunc func(http.ResponseWriter, *http.Request, *util.RequestContext)) (httpStatus int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(r.URL.Path, err, ";trace", string(debug.Stack()))
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}
	}()

	c := util.NewRequestContext()

	realFunc(w, r, c)
	httpStatus = http.StatusOK
	return
}

func startServe() {
	port := os.Args[1]

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      &handler{},
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func stop() {
	fmt.Println("stopping...")

	chSignal := make(chan os.Signal)
	signal.Notify(chSignal, os.Kill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSignal)

	fmt.Println("stopped.")
}
