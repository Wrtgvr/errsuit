package httpadap

import (
	"net/http"

	errsuit "github.com/wrtgvr/errsuit/core"
)

// Creates ErrorHandler and call HandleError
func HandleError(ctx HttpContext, err error, format errsuit.ResponseFormat) bool {
	return NewHttpErrorHandler(errsuit.Config{
		Format: format,
	}).HandleError(ctx, err)
}

func Handle(w http.ResponseWriter, r *http.Request, h *HttpErrorHandler, err error) bool {
	return h.HandleError(*ContextFromHttp(w, r), err)
}
