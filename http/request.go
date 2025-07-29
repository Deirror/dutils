package http

import (
	"net/http"
)

// GetPathParam returns the path parameter by key as a string.
// If the parameter does not exist, returns empty string.
func GetPathParam(r *http.Request, key string) string {
	return r.PathValue(key)
}

// GetQueryParam returns the first query parameter by key as string.
// Returns empty string if key does not exist.
func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
