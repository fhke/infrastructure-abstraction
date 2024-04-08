package request

import "errors"

var (
	ErrEmptyRepo = errors.New("repository is empty")
)

type BuildStack struct {
	Name        string   `json:"name"`
	Repository  string   `json:"repository"`
	ModuleNames []string `json:"moduleNames"`
}

func (b BuildStack) Validate() error {
	if b.Name == "" {
		return ErrEmptyName
	}
	if b.Repository == "" {
		return ErrEmptyRepo
	}
	return nil
}
