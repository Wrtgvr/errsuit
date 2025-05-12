package ginadap

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr2/errsuit"
)

// Gin error handler with `errsuit.Logger`.
type GinErrorHandler struct {
	logger *errsuit.Logger
}

// Returns `GinErrorHandler` with given `errsuit.Logger` (may be nil).
func NewGinErrorHandler(logger *errsuit.Logger) *GinErrorHandler {
	return &GinErrorHandler{
		logger: logger,
	}
}

// Send response via gin.Context with err HTTP status code, err message and err type (type e.g. `errsuit.TypeNotFound`).
// If err is type of `error` then converts it to `AppError`.
// Return `false` if err is nil, otherwise return true.
func (h *GinErrorHandler) HandleError(c *gin.Context, err error) bool {
	if err == nil || err.(*errsuit.AppError) == nil {
		return false
	}
	appErr := errsuit.AsAppError(err)

	if appErr.ShouldLog() {
		appErr.LogError()
	}

	msg := appErr.Message

	if appErr.Type == errsuit.TypeInternal {
		msg = "internal server error"
	}

	c.JSON(appErr.Code, gin.H{
		"error": msg,
		"type":  appErr.Type,
	})

	c.Abort()

	return true
}
