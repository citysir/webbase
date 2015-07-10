package app

import (
	"fmt"
	"net/http"
	"web/util"
	"web/webapi"
)

var pathRoute = map[string]func(http.ResponseWriter, *http.Request, *util.RequestContext){
	"/": webapi.Echo,
}

type handler struct {
}

func (this *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var httpStatus int
	if pathFunc, present := pathRoute[r.URL.Path]; present {
		httpStatus = safeRun(w, r, pathFunc)
	} else {
		httpStatus = http.StatusNotFound
		http.NotFound(w, r)
	}

	fmt.Println(r.URL.Path, httpStatus)
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

func startWebServe(port string) {
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
