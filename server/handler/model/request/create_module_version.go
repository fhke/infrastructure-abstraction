package request

import (
	"errors"

	"github.com/Masterminds/semver"
)

var (
	ErrEmptyVersion = errors.New("version is empty")
	ErrEmptySource  = errors.New("source is empty")
)

type CreateModuleVersion struct {
	Name    string          `json:"name"`
	Source  string          `json:"source"`
	Version *semver.Version `json:"version"`
}

func (c *CreateModuleVersion) Validate() error {
	if c.Name == "" {
		return ErrEmptyName
	}
	if c.Source == "" {
		return ErrEmptySource
	}
	if c.Version == nil {
		return ErrEmptyVersion
	}
	return nil
}
