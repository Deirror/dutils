package http

import (
	"net/http"

	"github.com/Deirror/dutils/cfg"
)

// NewStdServerFromConfig creates a standard HTTP server using the provided configuration and handler.
func NewStdServerFromConfig(cfg *cfg.ServerConfig, h http.Handler) *http.Server {
	return &http.Server{
		Addr:         cfg.Port,
		Handler:      h,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
