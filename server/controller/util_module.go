package controller

import (
	"fmt"

	"github.com/fhke/infrastructure-abstraction/server/storage/module/model"
)

func (c *controllerImpl) getLatestModuleVersion(module string) (model.ModuleVersion, error) {
	var latest model.ModuleVersion

	mvs, err := c.moduleRepo.GetVersions(module)
	if err != nil {
		return model.ModuleVersion{}, fmt.Errorf("error getting module versions: %w", err)
	}

	for _, mv := range mvs {
		if latest.Version == nil || mv.Version.GreaterThan(latest.Version) {
			latest = mv
		}
	}

	return latest, nil
}
