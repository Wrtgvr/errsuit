package ginadap

import (
	errsuit "github.com/wrtgvr/errsuit/core"
)

// `net/http` error handler.
// logger is optional. If logger is nil errors won't be logged.
type HttpErrorHandler struct {
	logger *errsuit.Logger
	cfg    errsuit.Config
}

// Returns `HttpErrorHandler` with given `errsuit.Logger` (may be nil).
func NewHttpErrorHandler(logger *errsuit.Logger) *HttpErrorHandler {
	return &HttpErrorHandler{
		logger: logger,
	}
}

// Send response via gin.Context with err HTTP status code, err message and err type (type e.g. `errsuit.TypeNotFound`).
// If err is type of `error` then converts it to `AppError`.
// Return `false` if err is nil, otherwise return true.
func (h *HttpErrorHandler) HandleError(ctx HttpContext, err error) bool {
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
