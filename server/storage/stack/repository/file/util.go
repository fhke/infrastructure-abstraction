package file

import (
	"github.com/fhke/infrastructure-abstraction/server/storage/stack/model"
)

func (f *fileRepository) findStack(name, repository string) (model.Stack, int) {
	for i, st := range f.f.Data.Stacks {
		if stackEq(name, repository, st) {
			return st, i
		}
	}

	return model.Stack{}, -1
}

func stackEq(name, repository string, st model.Stack) bool {
	return name == st.Name && repository == st.Repository
}
