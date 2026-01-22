package app

import "context"

type Runner interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}
