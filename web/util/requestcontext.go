package util

type RequestContext struct {
}

func NewRequestContext() *RequestContext {
	context := new(RequestContext)
	return context
}
