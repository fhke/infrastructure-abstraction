package deploy

import (
	"context"
	"fmt"

	"github.com/fhke/infrastructure-abstraction/sdk/executor"
	"github.com/fhke/infrastructure-abstraction/sdk/render"
)

func New[T any](ex executor.Executor[T], re render.Renderer[T]) Deployer[T] {
	return &runnerImpl[T]{
		re: re,
		ex: ex,
	}
}

func (r *runnerImpl[T]) Deploy(ctx context.Context, stack render.Stack) error {
	rendered, err := r.re.RenderStack(ctx, stack)
	if err != nil {
		return fmt.Errorf("error rendering stack: %w", err)
	}

	if err := r.ex.Exec(rendered); err != nil {
		return fmt.Errorf("error executing stack: %w", err)
	}

	return nil
}
