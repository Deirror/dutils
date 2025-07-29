package components

import (
	"net/http"

	"github.com/a-h/templ"
)

// Writes rendered templ component to the response.
func SendTempl(w http.ResponseWriter, r *http.Request, comp templ.Component) {
	w.Header().Set("Content-Type", "text/html")
	templ.Handler(comp).ServeHTTP(w, r)
}
