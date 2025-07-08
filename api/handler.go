package api

import (
	"context"
	"net/http"

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
