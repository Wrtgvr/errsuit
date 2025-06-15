package ginadap

import (
	"github.com/gin-gonic/gin"
	errsuit "github.com/wrtgvr/errsuit/core"
)

// Creates ErrorHandler and call HandleError
func HandleError(ctx GinContext, err error, format errsuit.ResponseFormat) bool {
	return NewGinErrorHandler(errsuit.Config{
		Format: format,
	}, nil).HandleError(ctx, err)
}

func Handle(ctx *gin.Context, h *GinErrorHandler, err error) bool {
	return h.HandleError(*ContextFromGin(ctx), err)
}
