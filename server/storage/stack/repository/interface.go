package repository

import (
	"context"

	"github.com/fhke/infrastructure-abstraction/server/storage/stack/model"
)

type Repository interface {
	GetStack(ctx context.Context, name, repository string) (model.Stack, error)
	AddStack(context.Context, model.Stack) error
	UpdateStack(context.Context, model.Stack) error
}
