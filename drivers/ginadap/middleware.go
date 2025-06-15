package ginadap

import (
	"github.com/gin-gonic/gin"
	errsuit "github.com/wrtgvr/errsuit/core"
)

const handlerFuncContextKey = "errsuit.hanalderfunc"

func InjectErrHandlerMiddleware(h *GinErrorHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		errCtx := ContextFromGin(c)

		fn := errsuit.MakeHandlerFunc(h, errCtx)

		c.Set(handlerFuncContextKey, fn)

		c.Next()
	}
}

func ErrHandlerFuncFromContext(c *gin.Context) func(error) bool {
	v, ok := c.Get(handlerFuncContextKey)
	if !ok {
		panic("errsuit: no error handler func in context")
	}
	fn, ok := v.(func(error) bool)
	if !ok {
		panic("errsuit: invalid handler func type in context")
	}

	return fn
}
