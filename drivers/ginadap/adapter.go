package ginadap

import (
	errsuit "github.com/wrtgvr/errsuit/core"
)

// Gin error handler.
// logger is optional. If logger is nil errors won't be logged.
type GinErrorHandler struct {
	cfg errsuit.Config
}

// Returns `GinErrorHandler` with given `errsuit.Logger` (may be nil).
func NewGinErrorHandler(cfg errsuit.Config) *GinErrorHandler {
	return &GinErrorHandler{
		cfg: cfg,
	}
}

// Send response via gin.Context with err HTTP status code, err message and err type (type e.g. `errsuit.TypeNotFound`).
// If err is type of `error` then converts it to `AppError`.
// Return `false` if err is nil, otherwise return true.
func (h GinErrorHandler) HandleError(ctx errsuit.Context, err error) bool {
	ginCtx, ok := ctx.(*GinContext)
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

	errsuit.WriteError(ginCtx, appErr, h.cfg.Format)
	ginCtx.C.Abort()

	return true
}
