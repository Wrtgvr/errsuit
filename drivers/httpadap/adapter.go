package httpadap

import (
	errsuit "github.com/wrtgvr/errsuit/core"
)

// `net/http` error handler.
// logger is optional. If logger is nil errors won't be logged.
type HttpErrorHandler struct {
	logger errsuit.ErrorLogger
	cfg    errsuit.Config
}

// Returns `HttpErrorHandler` with given `errsuit.Logger` (may be nil).
func NewHttpErrorHandler(cfg errsuit.Config, logger errsuit.ErrorLogger) *HttpErrorHandler {
	return &HttpErrorHandler{
		logger: logger,
		cfg:    cfg,
	}
}

// Send response via gin.Context with err HTTP status code, err message and err type (type e.g. `errsuit.TypeNotFound`).
// If err is type of `error` then converts it to `AppError`.
// Return `false` if err is nil, otherwise return true.
func (h *HttpErrorHandler) HandleError(ctx *HttpContext, err error) bool {
	appErr := errsuit.AsAppError(err)
	if appErr == nil {
		return false
	}

	if appErr.ShouldLog() && h.logger != nil {
		h.logger.LogError(appErr)
	}

	errsuit.WriteError(ctx, appErr, h.cfg.Format)

	return true
}

// Creates ErrorHandler and call HandleError
func HandleError(ctx *HttpContext, err error, format errsuit.ResponseFormat, logger errsuit.ErrorLogger) bool {
	return NewHttpErrorHandler(errsuit.Config{
		Format: format,
	}, logger).HandleError(ctx, err)
}
