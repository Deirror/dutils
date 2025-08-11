package api

import (
	"errors"
)

// ErrorResp is a simple error response struct used for API responses.
// `MsgKey` is a frontend-friendly message key for localization (i18n).
// `InternalMsg` is not serialized in JSON and is only used for logging.
type ErrorResp struct {
	Status      int    `json:"status"`      // HTTP status code
	MsgKey      string `json:"message_key"` // i18n key
	InternalMsg string `json:"-"`           // Server-side log message
}

// NewErrorResp constructs a new ErrorResp with both client-facing key and internal message.
func NewErrorResp(status int, msgKey, internalMsg string) *ErrorResp {
	return &ErrorResp{
		Status:      status,
		MsgKey:      msgKey,
		InternalMsg: internalMsg,
	}
}

// NewClientErrorResp constructs an ErrorResp with only a client-facing key.
func NewClientErrorResp(status int, msgKey string) *ErrorResp {
	return NewErrorResp(status, msgKey, "")
}

// Error implements the error interface and returns the MsgKey.
func (e *ErrorResp) Error() string {
	return e.MsgKey
}

// FromError attempts to convert a generic error into an *ErrorResp.
// Returns nil if the error is not of the correct type.
func FromError(err error) *ErrorResp {
	if err == nil {
		return nil
	}
	var errResp *ErrorResp
	if ok := AsErrorResp(err, &errResp); ok {
		return errResp
	}
	return nil
}

// AsErrorResp checks whether the error is of type *ErrorResp.
func AsErrorResp(err error, target **ErrorResp) bool {
	return errors.As(err, target)
}
