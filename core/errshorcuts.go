package errsuit

import "net/http"

// NewBadRequest returns a bad request error.
func NewBadRequest(msg string, err error, log bool) *AppError {
	return New(msg, http.StatusBadRequest, TypeBadRequest, err, log)
}

// NewNotFound returns a not found error.
func NewNotFound(msg string, err error, log bool) *AppError {
	return New(msg, http.StatusNotFound, TypeNotFound, err, log)
}

// NewInternal returns an internal error.
func NewInternal(msg string, err error, log bool) *AppError {
	return New(msg, http.StatusInternalServerError, TypeInternal, err, log)
}

// NewUnauthorized returns an unauthorized error.
func NewUnauthorized(msg string, err error, log bool) *AppError {
	return New(msg, http.StatusUnauthorized, TypeUnauthorized, err, log)
}

// NewForbidden returns an forbidden error.
func NewForbidden(msg string, err error, log bool) *AppError {
	return New(msg, http.StatusForbidden, TypeForbidden, err, log)
}

// NewConflict returns an conflict error.
func NewConflict(msg string, err error, log bool) *AppError {
	return New(msg, http.StatusConflict, TypeConflict, err, log)
}
