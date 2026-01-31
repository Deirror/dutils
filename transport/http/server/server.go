package httpsrv

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
)

type Server struct {
	log *slog.Logger

	srv *http.Server
}

func New(cfg *Config, log *slog.Logger, h http.Handler) *Server {
	return &Server{
		log: log,
		srv: NewStdServer(cfg, h),
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.log.Info("HTTP Server is starting", slog.String("addr", s.srv.Addr))

	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("HTTP Server is shutting down")

	return s.srv.Shutdown(ctx)
}

// New creates a standard HTTP server using the provided configuration and handler.
func NewStdServer(cfg *Config, h http.Handler) *http.Server {
	addr := cfg.Port
	if !strings.HasPrefix(cfg.Port, ":") {
		addr = ":" + cfg.Port
	}

	return &http.Server{
		Addr:         addr,
		Handler:      h,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
