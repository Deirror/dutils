package appx

import (
	"context"
	"log/slog"
	"sync"
)

type Runner interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type App struct {
	log *slog.Logger

	runners []Runner
}

func New(log *slog.Logger, runners ...Runner) *App {
	return &App{
		log:     log,
		runners: runners,
	}
}

func (a *App) Run(ctx context.Context) error {
	var wg sync.WaitGroup

	// shared error channel
	errCh := make(chan error, len(a.runners))

	a.log.Info("app starting")

	// start runners
	a.start(ctx, &wg, errCh)

	// wait for stop signal or runner error
	select {
	case <-ctx.Done():
		a.log.Info("app shutdown requested")
	case err := <-errCh:
		a.log.Error("app stopping due to runner error", slog.Any("err", err))
	}

	a.log.Info("app shutting down")

	// stop runners
	a.shutdown(ctx, &wg)

	wg.Wait()
	return nil
}

func (a *App) start(ctx context.Context, wg *sync.WaitGroup, errCh chan<- error) {
	for i, runner := range a.runners {
		i := i
		runner := runner

		wg.Go(func() {
			a.log.Info("starting", slog.Int("runner", i))
			if err := runner.Start(ctx); err != nil {
				a.log.Error("exited",
					slog.Int("runner", i),
					slog.Any("err", err),
				)
				errCh <- err
			}
		})
	}
}

func (a *App) shutdown(ctx context.Context, wg *sync.WaitGroup) {
	for i, runner := range a.runners {
		i := i
		runner := runner

		wg.Go(func() {
			a.log.Info("stopping", slog.Int("runner", i))
			if err := runner.Shutdown(ctx); err != nil {
				a.log.Error("shutdown failed",
					slog.Int("runner", i),
					slog.Any("err", err),
				)
			}
		})
	}
	wg.Wait()
}
