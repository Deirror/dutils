package http

import (
	"net/http"
	"strings"

	"github.com/Deirror/dutils/cfg"
)

// NewStdServerFromConfig creates a standard HTTP server using the provided configuration and handler.
func NewStdServerFromConfig(cfg *cfg.ServerConfig, h http.Handler) *http.Server {
	if !strings.HasPrefix(cfg.Port, ":") {
		cfg.Port = ":" + cfg.Port
	}

	return &http.Server{
		Addr:         cfg.Port,
		Handler:      h,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
