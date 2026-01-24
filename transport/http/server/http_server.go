package httpsrv

import (
	"net/http"
	"strings"
)

// New creates a standard HTTP server using the provided configuration and handler.
func New(cfg *cfg.ServerConfig, h http.Handler) *http.Server {
	addr := cfg.Port
	if !strings.HasPrefix(cfg.Port, ":") {
		addr = ":" + cfg.Port
	}

	return &http.Server{
		Addr:         cfg.Port,
		Handler:      h,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
