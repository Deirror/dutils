package api

import (
	"errors"
)

// Simple error response struct with json notation.
type ErrorResp struct {
	Status int    `json:"status"`  // HTTP status code with a json tag
	Msg    string `json:"message"` // Human-readable message with a json tag
}

func NewErrorResp(status int, msg string) *ErrorResp {
	return &ErrorResp{
		Status: status,
		Msg:    msg,
	}
}

// Implementation func of error interface.
func (e *ErrorResp) Error() string {
	return e.Msg
}

// Convertion from error to ErrorResp if possbile.
func FromError(err error) *ErrorResp {
	if err == nil {
		return nil
	}

	var errResp *ErrorResp
	if ok := AsErrorResp(err, &errResp); !ok {
		return nil
	}

	return errResp
}

// Helper func for checking if error is of type *ErroResp.
func AsErrorResp(err error, target **ErrorResp) bool {
	return errors.As(err, target)
}
