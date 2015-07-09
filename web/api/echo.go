package api

import (
	"../util"
	"fmt"
	"net/http"
)

func Echo(w http.ResponseWriter, r *http.Request, c *util.RequestContext) {
	fmt.Fprintf(w, "Hello, Gopher: %s\n", httputil.GetRemoteIP(r))
}
