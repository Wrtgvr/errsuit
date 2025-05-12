package errsuit

import (
	"fmt"
	"io"
)

// Error Logger, user to log errors
type ErrorLogger interface {
	LogError(error)
}

// Logger used to log errors with formatting them by format func.
type Logger struct {
	// Logs output.
	out io.Writer
	// Error format func.
	format func(error) string
}

// Creates logger with given out and default formatter.
func NewLogger(out io.Writer) *Logger {
	return &Logger{
		out:    out,
		format: defaultFormat,
	}
}

// Default error formatter.
func defaultFormat(err error) string {
	return fmt.Sprintf("[ERROR] %v", err)
}

// Log error formatted by `logger.format(err)` in `logger.out`.
func (l *Logger) LogError(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(l.out, l.format(err))
}
