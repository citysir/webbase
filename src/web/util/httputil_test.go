package util

import "testing"

func TestHttpGet(t *testing.T) {
	data, err := HttpGet("http://www.baidu.com")
	if err != nil {
		t.Error("test TestHttpGet failed")
	} else {
		t.Log("test TestHttpGet pass:", len(data))
	}
}
