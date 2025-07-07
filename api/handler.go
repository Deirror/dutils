package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// A handler func which accepts context field and returns error response.
type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request) *ErrorResp

// Custom type of error handler func mostly used for custom transport logic.
type ErrorHandlerFunc func(http.ResponseWriter, *http.Request, *ErrorResp)

// A wrapper func for handling errors from the called handler funcs.
// Uses custom func which handles error response.
// Defaults to DefaultErrorHandler if nothing is passed.
func Wrap(h HandlerFunc, onErr ErrorHandlerFunc) http.HandlerFunc {
	if onErr == nil {
		onErr = DefaultErrorHandler
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, ReqIDKey, uuid.NewString())
		if errResp := h(ctx, w, r); errResp != nil {
			onErr(w, r, errResp)
		}
	}
}

// Default func for json error handling.
func JSONErrorHandler(w http.ResponseWriter, r *http.Request, errResp *ErrorResp) {
	if errResp == nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"status":500,"message":"internal server error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errResp.Status)
	_ = json.NewEncoder(w).Encode(errResp)
}

// Default func for html error handling, with examplary html code.
func HTMLErrorHandler(w http.ResponseWriter, r *http.Request, errResp *ErrorResp) {
	if errResp == nil {
		errResp = NewErrorResp(http.StatusInternalServerError, "internal server error")
	}

	w.WriteHeader(errResp.Status)
	_, _ = w.Write([]byte(
		"<html><head><title>Error</title></head><body>" +
			"<h1>Error</h1>" +
			"<p>Status: " + http.StatusText(errResp.Status) + "</p>" +
			"<p>Message: " + errResp.Msg + "</p>" +
			"</body></html>",
	))
}

// Can be used in Wrap func as default one
func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, errResp *ErrorResp) {
	accept := r.Header.Get("Accept")
	if strings.Contains(accept, "text/html") {
		HTMLErrorHandler(w, r, errResp)
	} else {
		JSONErrorHandler(w, r, errResp)
	}
}
