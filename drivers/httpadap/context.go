package httpadap

import (
	"net/http"
)

type HttpContext struct {
	W http.ResponseWriter
	R *http.Request
}

func ContextFromHttp(w http.ResponseWriter, r *http.Request) *HttpContext {
	return &HttpContext{
		W: w,
		R: r,
	}
}

func (ctx HttpContext) SetHeader(key, value string) {
	ctx.W.Header().Add(key, value)
}

func (ctx HttpContext) GetHeader(key string) string {
	return ctx.R.Header.Get(key)
}

func (ctx HttpContext) SetStatus(code int) {
	ctx.W.WriteHeader(code)
}

func (ctx HttpContext) Write(body []byte) {
	ctx.W.Write(body)
}
