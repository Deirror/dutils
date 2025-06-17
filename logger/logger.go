package logger

import (
	"log/slog"
	"os"
)

// Inits the go standart logger, based on env mode.
func InitLogger(mode string) *slog.Logger {
	var h slog.Handler
	if mode == "dev" {
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	} else {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}
	logger := slog.New(h)
	return logger
}
