package repository

import "github.com/fhke/infrastructure-abstraction/server/storage/module/model"

type Repository interface {
	GetVersions(name string) ([]model.ModuleVersion, error)
	AddVersion(mv model.ModuleVersion) error
}
