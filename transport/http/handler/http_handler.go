package httphdl

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/Deirror/servette/trans/err"
)

// A handler func which accepts context field and returns an error.
type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request) *errx.Err

// Custom type of an error handler func mostly used for custom transport logic.
type ErrHandlerFunc func(context.Context, http.ResponseWriter, *http.Request, *errx.Err)

// A wrapper func for handling errors from the called handler funcs.
// Uses custom func which handles error.
// Defaults to DefaultErrHandler if nothing is passed.
func Wrap(h HandlerFunc, onErr ErrHandlerFunc) http.HandlerFunc {
	if onErr == nil {
		onErr = DefaultErrHandler
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, ReqIDKey, uuid.NewString())
		if errResp := h(ctx, w, r); errResp != nil {
			onErr(ctx, w, r, errResp)
		}
	}
}
