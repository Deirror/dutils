package json

import (
	"encoding/json"
	"io"
	"net/http"
)

// A wrapper for json data encoding.
func EncodeJSON(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(v)
}

// A func that can be used in handlers to write JSON.
// Includes status code and structured data to marshall.
func WriteJSON(w http.ResponseWriter, s int, v any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(s)
	return EncodeJSON(w, v)
}

// Can be called when already having a var of type T.
// Unmarshals data into the target.
func ParseJSONInto[T any](r io.Reader, target *T) error {
	return json.NewDecoder(r).Decode(target)
}

// Same as func ParseJSONInto, but initziliazes new var of T.
func ParseJSON[T any](r io.Reader) (T, error) {
	var v T
	err := json.NewDecoder(r).Decode(&v)
	return v, err
}
