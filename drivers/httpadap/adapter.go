package httpadap

import (
	"encoding/json"
	"net/http"

	"github.com/wrtgvr2/errsuit"
)

// `net/http` error handler.
// logger is optional. If logger is nil errors won't be logged.
type HttpErrorHandler struct {
	logger *errsuit.Logger
}

// Returns `HttpErrorHandler` with given `errsuit.Logger` (may be nil).
func NewHttpErrorHandler(logger *errsuit.Logger) *HttpErrorHandler {
	return &HttpErrorHandler{
		logger: logger,
	}
}

// Send response via `http.ResponseWriter` with err HTTP status code, err message and err type (type e.g. `errsuit.TypeNotFound`).
// If err is type of `error` then converts it to `AppError`.
// Return `false` if err is nil, otherwise return true.
func (h *HttpErrorHandler) HandleError(w http.ResponseWriter, err error) bool {
	appErr := errsuit.AsAppError(err)
	if appErr == nil {
		return false
	}

	if appErr.ShouldLog() {
		if h.logger != nil {
			h.logger.LogError(appErr)
		}
	}

	msg := appErr.Message

	if appErr.Type == errsuit.TypeInternal {
		msg = "internal server error"
	}

	w.WriteHeader(appErr.Code)
	json.NewEncoder(w).Encode(errsuit.ErrorResponse{
		ErrMsg: msg,
		Typ:    appErr.Type,
	})

	return true
}
