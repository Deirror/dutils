package api

import (
	"errors"
)

// Simple error response struct with json notation.
type ErrorResp struct {
	Status    int    `json:"status"`         // HTTP status code with a json tag
	ClientMsg string `json:"client_message"` // Human-readable message with a json tag
	ServerMsg string `json:"-"`              // Server side message, for logging purposes
}

func NewErrorResp(status int, clientMsg, serverMsg string) *ErrorResp {
	return &ErrorResp{
		Status:    status,
		ClientMsg: clientMsg,
		ServerMsg: serverMsg,
	}
}

// Implementation func of error interface.
func (e *ErrorResp) Error() string {
	return e.ClientMsg
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
