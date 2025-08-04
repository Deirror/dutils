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

// LogFunc logs a function call with optional extra attributes.
// It records the request ID from context and any error passed.
// Timing is NOT recorded here.
// Use this when you only want to log function entry or results without measuring duration.
func LogFunc(
	ctx context.Context,
	logger *slog.Logger,
	funcName string,
	err error,
	extraAttrs ...slog.Attr,
) {
	attrs := []slog.Attr{
		slog.Any(api.ReqIDKey, ctx.Value(api.ReqIDKey)),
		slog.Any("error", err),
	}
	attrs = append(attrs, extraAttrs...)

	level := slog.LevelInfo
	if err != nil {
		level = slog.LevelError
	}

	logger.LogAttrs(ctx, level, funcName, attrs...)
}

// LogFuncWithTiming logs a function call including the elapsed time since `begin`.
// It records the request ID from context, the duration, any error, and optional extra attributes.
// Use this when you want to measure and log the time taken by a function.
func LogFuncWithTiming(
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
