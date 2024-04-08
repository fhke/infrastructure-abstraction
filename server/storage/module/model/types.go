package model

import (
	"github.com/Masterminds/semver"
)

type (
	ModuleVersion struct {
		Name    string          `json:"name"`    // name is the user-facing name of the module.
		Source  string          `json:"source"`  // source is the terraform-readable source of the module.
		Version *semver.Version `json:"version"` // version is the version of the module.
	}
)
