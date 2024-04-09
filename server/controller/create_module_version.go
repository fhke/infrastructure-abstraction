package controller

import (
	"errors"
	"fmt"

	storageErr "github.com/fhke/infrastructure-abstraction/server/storage/common/errors"
	"github.com/fhke/infrastructure-abstraction/server/storage/module/model"
)

func (c *controllerImpl) CreateModuleVersion(module, source, version string) error {
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
