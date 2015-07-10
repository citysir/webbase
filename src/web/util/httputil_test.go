package util

import "testing"

func Test_HttpGet(t *testing.T) {
	data, err := HttpGet("http://www.baidu.com")
	if err != nil {
		t.Error("test TestHttpGet failed", err)
	} else {
		t.Log("test TestHttpGet pass:", len(data))
	}
}
