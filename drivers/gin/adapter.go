package ginadapter

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr2/errsuit"
)

func HandleError(c *gin.Context, err error) {
	appErr := errsuit.AsAppError(err)

	if appErr.ShouldLog() {
		appErr.LogError()
	}

	c.JSON(appErr.Code, gin.H{
		"error": appErr.Message,
		"type":  appErr.Type,
	})
}
