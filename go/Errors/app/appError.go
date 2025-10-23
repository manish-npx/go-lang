package app

import "fmt"

type AppError struct {
	Code    string // internal code like "NO_TEA", "BROKEN_KETTLE"
	Message string // message for logging or API response
	Err     error  // wrapped error (root cause)
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s : %v", e.Code, e.Message, e.Err)
}

// helper for wrapping
func WrapError(code string, msg string, err error) *AppError {
	return &AppError{Code: code, Message: msg, Err: err}
}
