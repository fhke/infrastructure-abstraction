package controller

import (
	"context"
	"errors"
	"fmt"

	storageErr "github.com/fhke/infrastructure-abstraction/server/storage/common/errors"
	moduleModel "github.com/fhke/infrastructure-abstraction/server/storage/module/model"
	stackModel "github.com/fhke/infrastructure-abstraction/server/storage/stack/model"
)

func (c *controllerImpl) SetStackModules(ctx context.Context, name, repo string, moduleVersions map[string]string) (stackModel.Stack, error) {
	log := c.log.With("operation", "set_stack_modules", "stack", map[string]string{"name": name, "repository": repo})
	st, err := c.stackRepo.GetStack(ctx, name, repo)
	if err != nil {
		log.Warnw("Error getting stack", "error", err)
		if errors.Is(err, storageErr.ErrNotFound) {
			return stackModel.Stack{}, ErrStackNotFound{Stack: name, Repo: repo}
		}
		return stackModel.Stack{}, fmt.Errorf("")
	}

	for modName, modVersion := range moduleVersions {
		mod, err := c.getModuleVersion(ctx, modName, modVersion)
		if err != nil {
			return stackModel.Stack{}, fmt.Errorf("error getting version %s for module %s: %w", modVersion, modName, err)
		}
		st.Modules[modName] = stackModel.Module{
			Source:  mod.Source,
			Version: mod.Version,
		}
	}

	if err := c.stackRepo.UpdateStack(ctx, st); err != nil {
		log.Errorw("Error updating stack", "error", err)
		return stackModel.Stack{}, fmt.Errorf("error updating stack: %w", err)
	}

	return st, nil
}

func (c *controllerImpl) getModuleVersion(ctx context.Context, name, version string) (moduleModel.ModuleVersion, error) {
	modVersions, err := c.moduleRepo.GetVersions(ctx, name)
	if err != nil {
		return moduleModel.ModuleVersion{}, fmt.Errorf("error getting module %q: %w", name, err)
	}

	for _, modVersion := range modVersions {
		if modVersion.Version == version {
			return modVersion, nil
		}
	}

	return moduleModel.ModuleVersion{}, ErrModuleVersionNotFound
}
