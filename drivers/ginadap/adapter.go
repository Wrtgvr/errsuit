package ginadap

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr2/errsuit"
)

// Send response via gin.Context with err HTTP status code, err message and err type (type e.g. `errsuit.TypeNotFound`).
// If err is type of `error` then converts it to `AppError`.
// Return `false` if err is nil, otherwise return true.
func HandleError(c *gin.Context, err error) bool {
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
