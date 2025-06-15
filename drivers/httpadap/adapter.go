package httpadap

import (
	errsuit "github.com/wrtgvr/errsuit/core"
)

// `net/http` error handler.
// logger is optional. If logger is nil errors won't be logged.
type HttpErrorHandler struct {
	cfg errsuit.Config
}

// Returns `HttpErrorHandler` with given `errsuit.Logger` (may be nil).
func NewHttpErrorHandler(cfg errsuit.Config) *HttpErrorHandler {
	return &HttpErrorHandler{
		cfg: cfg,
	}
}

// Send response via gin.Context with err HTTP status code, err message and err type (type e.g. `errsuit.TypeNotFound`).
// If err is type of `error` then converts it to `AppError`.
// Return `false` if err is nil, otherwise return true.
func (h HttpErrorHandler) HandleError(ctx errsuit.Context, err error) bool {
	httpCtx, ok := ctx.(*HttpContext)
	if !ok {
		panic("GinErrorHandler: invlaid context type passed")
	}

	appErr := errsuit.AsAppError(err)
	if appErr == nil {
		return false
	}

	if appErr.ShouldLog() && h.cfg.Logger != nil {
		h.cfg.Logger.LogError(appErr)
	}

	errsuit.WriteError(httpCtx, appErr, h.cfg.Format)

	return true
}
