package request

import (
	"errors"

	"github.com/Masterminds/semver"
)

var (
	ErrEmptyVersion = errors.New("version is empty")
	ErrEmptySource  = errors.New("source is empty")
	ErrNotSemver    = errors.New("version is not a semver")
)

type CreateModuleVersion struct {
	Name    string `json:"name"`
	Source  string `json:"source"`
	Version string `json:"version"`
}

func (c *CreateModuleVersion) Validate() error {
	if c.Name == "" {
		return ErrEmptyName
	}
	if c.Source == "" {
		return ErrEmptySource
	}
	if _, err := semver.NewVersion(c.Version); err != nil {
		return ErrEmptyVersion
	}
	return nil
}
