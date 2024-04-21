package controller

import (
	"context"
	"errors"
	"fmt"

	storageErr "github.com/fhke/infrastructure-abstraction/server/storage/common/errors"
	"github.com/fhke/infrastructure-abstraction/server/storage/stack/model"
	"k8s.io/apimachinery/pkg/util/sets"
)

func (c *controllerImpl) BuildStack(ctx context.Context, name, repo string, moduleNames sets.Set[string]) (model.Stack, error) {
	log := c.log.With("operation", "build_stack", "stack", map[string]string{"name": name, "repository": repo})

	// Get existing stack from DB
	isCreate := false
	st, err := c.stackRepo.GetStack(ctx, name, repo)
	if err != nil {
		if errors.Is(err, storageErr.ErrNotFound) {
			log.Debug("Building new stack")
			isCreate = true
			st = model.Stack{
				Name:       name,
				Repository: repo,
				Modules:    make(map[string]model.Module),
			}
		} else {
			log.Errorw("Error getting stack", "error", err)
			return model.Stack{}, fmt.Errorf("error getting stack: %w", err)
		}
	} else {
		log.Debug("Updating existing stack")
	}

	// indicates whether persisted data has changed
	hasChanges := false

	// Update Stack with missing module names
	for modName := range moduleNames {
		if _, ok := st.Modules[modName]; ok {
			// Module already exists in stack, skip it
			continue
		}
		hasChanges = true
		// get latest version of module
		latestMv, err := c.getLatestModuleVersion(ctx, modName)
		if err != nil {
			if errors.Is(err, storageErr.ErrNotFound) {
				// request specified a nonexistent module
				log.Warnw("Found nonexistent module in stack", "module", modName)
				return model.Stack{}, ErrModuleNotFound{Name: modName}
			}
			log.Errorw("Error getting module", "error", err)
			return model.Stack{}, fmt.Errorf("error getting module: %w", err)
		}
		log.Debugw("Adding module", "module", modName, "version", latestMv.Version, "source", latestMv.Source)
		st.Modules[modName] = model.Module{
			Source:  latestMv.Source,
			Version: latestMv.Version,
		}
	}

	// persist updated stack
	if isCreate {
		if err := c.stackRepo.AddStack(ctx, st); err != nil {
			log.Errorw("Error creating stack", "error", err)
			return model.Stack{}, fmt.Errorf("error creating stack: %w", err)
		}
		log.Debug("Created stack")
	} else if hasChanges {
		if err := c.stackRepo.UpdateStack(ctx, st); err != nil {
			log.Errorw("Error updating stack", "error", err)
			return model.Stack{}, fmt.Errorf("error updating stack: %w", err)
		}
		log.Debug("Updated stack")
	} else {
		log.Debug("No changes to persist")
	}

	return st, nil
}
