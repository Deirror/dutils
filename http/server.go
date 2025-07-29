package http

import (
	"net/http"

	cfg "github.com/Deirror/dutils/cfg/http"
)

// A wrapper func for creating a http server with a config and handler/router.
func NewHTTPServerWithConfig(cfg *cfg.HTTPServerConfig, h http.Handler) *http.Server {
	return &http.Server{
		Addr:         cfg.Port,
		Handler:      h,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
