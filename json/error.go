package json

import "net/http"

// Simple error response struct with json notation.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Sends structured json error.
// As marshal can fail, it returns error.
func SendErrorJSON(w http.ResponseWriter, status int, msg string) error {
	return WriteJSON(w, status, ErrorResponse{Error: msg})
}
