package ginadap

import (
	"github.com/gin-gonic/gin"
	errsuit "github.com/wrtgvr2/errsuit/core"
)

// Gin error handler.
// logger is optional. If logger is nil errors won't be logged.
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
	appErr := errsuit.AsAppError(err)
	if appErr == nil {
		return false
	}

	if appErr.ShouldLog() {
		if h.logger != nil {
			h.logger.LogError(appErr)
		}
	}

	c.JSON(appErr.Code, errsuit.BuildErrorResp(appErr))

	c.Abort()

	return true
}
