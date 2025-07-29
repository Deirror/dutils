package components

import (
	"net/http"

	"github.com/a-h/templ"
)

// Writes rendered templ component to the response.
func SendTempl(w http.ResponseWriter, r *http.Request, comp templ.Component) {
	templ.Handler(comp).ServeHTTP(w, r)
}
