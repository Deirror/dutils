package logger

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/Deirror/dutils/api"
	"github.com/Deirror/dutils/env"
)

// Inits the go standart logger, based on env mode.
func InitLogger(mode string) *slog.Logger {
	var h slog.Handler
	if mode == env.Dev {
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	} else {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}
	logger := slog.New(h)
	return logger
}

// Wrapper func, used for dev/prod-ready logging of a service.
func LogService(
	ctx context.Context,
	logger *slog.Logger,
	funcName string,
	begin time.Time,
	err error,
	extraAttrs ...slog.Attr,
) {
	attrs := []slog.Attr{
		slog.Any(api.ReqIDKey, ctx.Value(api.ReqIDKey)),
		slog.Duration("took", time.Since(begin)),
		slog.Any("error", err),
	}
	attrs = append(attrs, extraAttrs...)

	level := slog.LevelInfo
	if err != nil {
		level = slog.LevelError
	}

	logger.LogAttrs(ctx, level, funcName, attrs...)
}

// Wrapper func, used for logging handler func errors.
func LogHandlerError(
	ctx context.Context,
	logger *slog.Logger,
	path string,
	err error,
	extraAttrs ...slog.Attr,
) {
	if err == nil {
		return
	}

	attrs := []slog.Attr{
		slog.Any(api.ReqIDKey, ctx.Value(api.ReqIDKey)),
		slog.Any("error", err),
	}
	attrs = append(attrs, extraAttrs...)

	logger.LogAttrs(ctx, slog.LevelError, path, attrs...)
}
