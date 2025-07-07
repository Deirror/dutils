package api

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
