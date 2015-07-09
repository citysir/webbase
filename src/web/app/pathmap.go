package app

import (
	"../api"
	"../util"
	"net/http"
)

var pathMap = map[string]func(http.ResponseWriter, *http.Request, *util.RequestContext){
	"/": api.Echo,
}
