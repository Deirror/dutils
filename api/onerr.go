package api

import (
	"net/http"
	"strings"

	djson "github.com/Deirror/dutils/json"
)

// Can be used in Wrap func as default one
func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, errResp *ErrorResp) {
	accept := r.Header.Get("Accept")
	if strings.Contains(accept, "text/html") {
		HTMLErrorHandler(w, r, errResp)
	} else {
		JSONErrorHandler(w, r, errResp)
	}
}

// Default func for json error handling.
func JSONErrorHandler(w http.ResponseWriter, r *http.Request, errResp *ErrorResp) {
	if errResp == nil {
		errResp = NewErrorResp(http.StatusInternalServerError, "internal server error", "errResp is nil")
	}
	djson.WriteJSON(w, errResp.Status, errResp)
}

// Default func for html error handling, with examplary html code.
func HTMLErrorHandler(w http.ResponseWriter, r *http.Request, errResp *ErrorResp) {
	if errResp == nil {
		errResp = NewErrorResp(http.StatusInternalServerError, "internal server error", "errResp is nil")
	}

	w.WriteHeader(errResp.Status)
	_, _ = w.Write([]byte(
		"<html><head><title>Error</title></head><body>" +
			"<h1>Error</h1>" +
			"<p>Status: " + http.StatusText(errResp.Status) + "</p>" +
			"<p>Message: " + errResp.ClientMsg + "</p>" +
			"</body></html>",
	))
}
