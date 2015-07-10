package webapi

import (
	"fmt"
	"net/http"
	"web/util"
)

func Echo(w http.ResponseWriter, r *http.Request, c *util.RequestContext) {
	fmt.Fprintf(w, "Hello, Gopher: %s\n", util.GetRemoteIP(r))
}
