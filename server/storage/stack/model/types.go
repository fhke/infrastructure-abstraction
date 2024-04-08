package model

import (
	"github.com/Masterminds/semver"
	"github.com/fhke/infrastructure-abstraction/server/storage/common/concurrency"
)

type (
	Stack struct {
		concurrency.OptimisticLock `json:",inline"`

		Name       string            `json:"name"`
		Repository string            `json:"repository"`
		Modules    map[string]Module `json:"modules"`
	}

	Module struct {
		Source  string          `json:"source"`
		Version *semver.Version `json:"version"`
	}
)
