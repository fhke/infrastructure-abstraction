package controller

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/fhke/infrastructure-abstraction/server/storage/module/model"
)

func (c *controllerImpl) getLatestModuleVersion(ctx context.Context, module string) (model.ModuleVersion, error) {
	var (
		latest        model.ModuleVersion
		latestVersion *semver.Version
	)

	mvs, err := c.moduleRepo.GetVersions(ctx, module)
	if err != nil {
		return model.ModuleVersion{}, fmt.Errorf("error getting module versions: %w", err)
	}

	for i, mv := range mvs {
		thisVersion, err := semver.NewVersion(mv.Version)
		if err != nil {
			return model.ModuleVersion{}, fmt.Errorf("error parsing semver %q for version %d: %w", mv.Version, i, err)
		}
		if latestVersion == nil || thisVersion.GreaterThan(latestVersion) {
			latest = mv
			latestVersion = thisVersion
		}
	}

	return latest, nil
}
