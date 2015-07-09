package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIp       = "X-Real-Ip"
)

var isRemoteAddrFirst bool

func SetRemoteAddrFirst(b bool) {
	isRemoteAddrFirst = b
}

func GetRemoteIP(r *http.Request) (remoteIP string) {
	if isRemoteAddrFirst {
		remoteIP = GetRemoteAddr(r)
		if remoteIP == "" {
			remoteIP = GetXForwardedFor(r)
		}
	} else {
		remoteIP = GetHeaderIP(r, XRealIp)
		if remoteIP == "" {
			remoteIP = GetRemoteAddr(r)
		}
	}

	return
}

func GetRemoteAddr(r *http.Request) (remoteIP string) {
	remoteAddr := strings.Split(r.RemoteAddr, ":")
	return remoteAddr[0]
}

func GetXForwardedFor(r *http.Request) (remoteIP string) {
	remoteIP = r.Header.Get(XForwardedFor)
	if remoteIP != "" { //可能出现这种 117.169.143.20, 120.203.215.3
		remoteIPS := strings.Split(remoteIP, ",")
		remoteIP = strings.TrimSpace(remoteIPS[len(remoteIPS)-1])
	}
	return
}

func GetHeaderIP(r *http.Request, header string) (remoteIP string) {
	remoteIP = r.Header.Get(header)
	if remoteIP != "" { //可能出现这种 117.169.143.20, 120.203.215.3
		remoteIPS := strings.Split(remoteIP, ",")
		remoteIP = strings.TrimSpace(remoteIPS[len(remoteIPS)-1])
	}
	return
}

func HttpGet(urlStr string) ([]byte, error) {
	r, _ := http.NewRequest("GET", urlStr, nil)
	return doRequest(r)
}

func HttpPut(urlStr string, data []byte) ([]byte, error) {
	r, _ := http.NewRequest("PUT", urlStr, bytes.NewReader(data))
	return doRequest(r)
}

func HttpPost(urlStr string, data url.Values) ([]byte, error) {
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	return doRequest(r)
}

func doRequest(request *http.Request) ([]byte, error) {
	client := &http.Client{Timeout: time.Duration(15 * time.Second)}

	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
