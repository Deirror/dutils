package logger

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/Deirror/dutils/api"
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

// Wrapper func, used for dev/prod-ready logging - based on the slog.LevelInfo
func LogWithAttrs(ctx context.Context, logger *slog.Logger, funcName string, begin time.Time, err error, extraAttrs ...slog.Attr) {
	attrs := []slog.Attr{
		slog.Any(api.ReqIDKey, ctx.Value(api.ReqIDKey)),
	}
	attrs = append(attrs, extraAttrs...)
	attrs = append(attrs, slog.Duration("took", time.Since(begin)))
	attrs = append(attrs, slog.Any("error", err))

	logger.LogAttrs(ctx, slog.LevelInfo, funcName, attrs...)
}
