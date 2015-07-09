package app

import (
	"net/http"
	"web/api"
	"web/util"
)

var pathMap = map[string]func(http.ResponseWriter, *http.Request, *util.RequestContext){
	"/": api.Echo,
}
