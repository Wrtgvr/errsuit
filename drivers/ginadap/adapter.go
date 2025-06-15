package ginadap

import (
	errsuit "github.com/wrtgvr/errsuit/core"
)

// Gin error handler.
// logger is optional. If logger is nil errors won't be logged.
type GinErrorHandler struct {
	logger errsuit.ErrorLogger
	cfg    errsuit.Config
}

// Returns `GinErrorHandler` with given `errsuit.Logger` (may be nil).
func NewGinErrorHandler(cfg errsuit.Config, logger errsuit.ErrorLogger) *GinErrorHandler {
	return &GinErrorHandler{
		logger: logger,
		cfg:    cfg,
	}
}

// Send response via gin.Context with err HTTP status code, err message and err type (type e.g. `errsuit.TypeNotFound`).
// If err is type of `error` then converts it to `AppError`.
// Return `false` if err is nil, otherwise return true.
func (h *GinErrorHandler) HandleError(ctx GinContext, err error) bool {
	appErr := errsuit.AsAppError(err)
	if appErr == nil {
		return false
	}

	if appErr.ShouldLog() && h.logger != nil {
		h.logger.LogError(appErr)
	}

	errsuit.WriteError(ctx, appErr, h.cfg.Format)
	ctx.C.Abort()

	return true
}
