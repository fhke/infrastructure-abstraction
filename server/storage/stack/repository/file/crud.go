package file

import (
	"os"
	"sync"

	"github.com/fhke/infrastructure-abstraction/server/storage/common/errors"
	"github.com/fhke/infrastructure-abstraction/server/storage/common/repository/file"
	"github.com/fhke/infrastructure-abstraction/server/storage/stack/model"
	"github.com/fhke/infrastructure-abstraction/server/storage/stack/repository"
)

func New(f *os.File) (repository.Repository, error) {
	fr, err := file.New[data](f)
	if err != nil {
		return nil, err
	}

	return &fileRepository{
		mu: new(sync.RWMutex),
		f:  fr,
	}, nil

}

func (f *fileRepository) GetStack(name, repository string) (model.Stack, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	if st, pos := f.findStack(name, repository); pos >= 0 {
		return st, nil
	}

	return model.Stack{}, errors.ErrNotFound
}

func (f *fileRepository) AddStack(st model.Stack) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if _, pos := f.findStack(st.Name, st.Repository); pos >= 0 {
		return errors.ErrAlreadyExists
	}

	f.f.Data.Stacks = append(f.f.Data.Stacks, st)

	return f.f.WriteData()
}

func (f *fileRepository) UpdateStack(st model.Stack) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	_, pos := f.findStack(st.Name, st.Repository)
	if pos == -1 {
		return errors.ErrNotFound
	}

	f.f.Data.Stacks[pos] = st

	return nil
}
