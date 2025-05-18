package errsuit

import (
	"errors"
	"fmt"
	"net/http"
)

// AppError represents a structured application error with optional logging and HTTP compatibility.
type AppError struct {
	// Type categorizes the error (e.g., "not_found", "internal").
	Type string `json:"type"`
	// Message is the error message shown to clients.
	Message string `json:"message"`
	// Code is the HTTP status code returned in responses.
	Code int `json:"code"`
	// Err is the original internal error (not included in JSON response).
	Err error `json:"-"`
	// Log indicates whether the error should be logged automatically.
	Log bool `json:"-"`
}

// Error returns a formatted string for the error.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// ShouldLog tells whether the error should be logged by a driver.
func (e *AppError) ShouldLog() bool {
	return e.Log
}

// New creates a new AppError.
func New(msg string, code int, typ string, err error, log bool) *AppError {
	if err == nil {
		err = errors.New(msg)
	}
	return &AppError{
		Type:    typ,
		Message: msg,
		Code:    code,
		Err:     err,
		Log:     log,
	}
}

// AsAppError converts any error into AppError.
func AsAppError(err error) *AppError {
	if err == nil {
		return nil
	}
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	return New("internal error", http.StatusInternalServerError, TypeInternal, err, true)
}

// Error types
const (
	TypeBadRequest   = "bad_request"
	TypeNotFound     = "not_found"
	TypeInternal     = "internal"
	TypeUnauthorized = "unauthorized"
	TypeForbidden    = "forbidden"
	TypeConflict     = "conflict"
)
