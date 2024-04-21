package repository

import (
	"context"

	"github.com/fhke/infrastructure-abstraction/server/storage/module/model"
)

type Repository interface {
	GetModuleNames(ctx context.Context) ([]string, error)
	GetVersions(ctx context.Context, name string) ([]model.ModuleVersion, error)
	AddVersion(ctx context.Context, mv model.ModuleVersion) error
}
