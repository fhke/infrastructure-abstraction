package file

import (
	"sync"

	"github.com/fhke/infrastructure-abstraction/server/storage/common/repository/file"
	"github.com/fhke/infrastructure-abstraction/server/storage/module/model"
)

type (
	// fileRepository is a concrete implementation of the storage/module/repository.Repository
	// interface, using a local file for storage.
	fileRepository struct {
		mu *sync.RWMutex

		f *file.FileRepo[data]
	}

	data struct {
		Modules map[string][]model.ModuleVersion `json:"modules"`
	}
)
