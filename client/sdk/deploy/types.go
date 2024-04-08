package deploy

import (
	"context"

	"github.com/fhke/infrastructure-abstraction/sdk/executor"
	"github.com/fhke/infrastructure-abstraction/sdk/render"
)

type (
	Deployer[T any] interface {
		Deploy(ctx context.Context, stack render.Stack) error
	}
	runnerImpl[T any] struct {
		re render.Renderer[T]
		ex executor.Executor[T]
	}
)
