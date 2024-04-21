package controller

import (
	"context"

	moduleRepository "github.com/fhke/infrastructure-abstraction/server/storage/module/repository"
	"github.com/fhke/infrastructure-abstraction/server/storage/stack/model"
	stackRepository "github.com/fhke/infrastructure-abstraction/server/storage/stack/repository"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/sets"
)

type (
	Controller interface {
		BuildStack(ctx context.Context, name, repo string, moduleNames sets.Set[string]) (model.Stack, error)
		SetStackModules(ctx context.Context, name, repo string, moduleVersions map[string]string) (model.Stack, error)
		CreateModuleVersion(ctx context.Context, module, source, version string) error
		GetModuleNames(ctx context.Context) ([]string, error)
	}

	controllerImpl struct {
		log *zap.SugaredLogger

		stackRepo  stackRepository.Repository
		moduleRepo moduleRepository.Repository
	}
)
