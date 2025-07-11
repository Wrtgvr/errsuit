package ginadap

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	C *gin.Context
}

func ContextFromGin(c *gin.Context) *GinContext {
	return &GinContext{
		C: c,
	}
}

func (ctx GinContext) SetHeader(key, value string) {
	ctx.C.Header(key, value)
}

func (ctx GinContext) GetHeader(key string) string {
	return ctx.C.GetHeader(key)
}

func (ctx GinContext) SetStatus(code int) {
	ctx.C.Status(code)
}

func (ctx GinContext) Write(body []byte) {
	ctx.C.Writer.Write(body)
}
