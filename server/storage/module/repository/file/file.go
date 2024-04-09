package file

import (
	"os"
	"sync"

	"github.com/fhke/infrastructure-abstraction/server/storage/common/errors"
	"github.com/fhke/infrastructure-abstraction/server/storage/common/repository/file"
	"github.com/fhke/infrastructure-abstraction/server/storage/module/model"
	"github.com/fhke/infrastructure-abstraction/server/storage/module/repository"
)

func New(f *os.File) (repository.Repository, error) {
	fr, err := file.New[data](f)
	if err != nil {
		return nil, err
	}

	if fr.Data.Modules == nil {
		fr.Data.Modules = make(map[string][]model.ModuleVersion)
	}

	return &fileRepository{
		mu: new(sync.RWMutex),
		f:  fr,
	}, nil

}

func (f *fileRepository) GetVersions(name string) ([]model.ModuleVersion, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	mvs := f.f.Data.Modules[name]
	if len(mvs) == 0 {
		return nil, errors.ErrNotFound
	}

	return mvs, nil
}

func (f *fileRepository) AddVersion(mv model.ModuleVersion) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if _, ok := f.f.Data.Modules[mv.Name]; !ok {
		f.f.Data.Modules[mv.Name] = make([]model.ModuleVersion, 0)
	}

	for _, exMv := range f.f.Data.Modules[mv.Name] {
		if exMv.Version == mv.Version {
			return errors.ErrAlreadyExists
		}
	}

	f.f.Data.Modules[mv.Name] = append(f.f.Data.Modules[mv.Name], mv)

	return f.f.WriteData()
}
