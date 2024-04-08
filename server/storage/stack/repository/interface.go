package repository

import "github.com/fhke/infrastructure-abstraction/server/storage/stack/model"

type Repository interface {
	GetStack(name, repository string) (model.Stack, error)
	AddStack(model.Stack) error
	UpdateStack(model.Stack) error
}
