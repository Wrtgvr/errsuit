package errsuit

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type AppError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Err     error  `json:"-"`
	Log     bool   `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) ShouldLog() bool {
	return e.Log
}

func (e *AppError) LogError() {
	log.Print(e.Error())
}

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

func AsAppError(err error) *AppError {
	if err == nil {
		return nil
	}
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	return New("internal error", http.StatusInternalServerError, TypeInternal, err, true)
}

const (
	TypeBadRequest = "bad_request"
	TypeNotFound   = "not_found"
	TypeInternal   = "internal"
)

// Shortcuts
func NewBadRequest(msg string, err error) *AppError {
	return New(msg, http.StatusBadRequest, TypeBadRequest, err, false)
}

func NewNotFound(msg string, err error) *AppError {
	return New(msg, http.StatusNotFound, TypeNotFound, err, false)
}

func NewInternal(msg string, err error) *AppError {
	return New(msg, http.StatusInternalServerError, TypeInternal, err, true)
}
