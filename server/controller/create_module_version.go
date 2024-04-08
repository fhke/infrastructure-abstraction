package controller

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver"
	storageErr "github.com/fhke/infrastructure-abstraction/server/storage/common/errors"
	"github.com/fhke/infrastructure-abstraction/server/storage/module/model"
)

func (c *controllerImpl) CreateModuleVersion(module, source string, version *semver.Version) error {
	if err := c.moduleRepo.AddVersion(model.ModuleVersion{
		Name:    module,
		Source:  source,
		Version: version,
	}); err != nil {
		if errors.Is(err, storageErr.ErrAlreadyExists) {
			return ErrModuleVersionAlreadyExists{Name: module, Version: version}
		}
		return fmt.Errorf("error adding module version: %w", err)
	}
	return nil
}
