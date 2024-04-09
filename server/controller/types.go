package controller

import (
	moduleRepository "github.com/fhke/infrastructure-abstraction/server/storage/module/repository"
	"github.com/fhke/infrastructure-abstraction/server/storage/stack/model"
	stackRepository "github.com/fhke/infrastructure-abstraction/server/storage/stack/repository"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/sets"
)

type (
	Controller interface {
		BuildStack(name, repo string, moduleNames sets.Set[string]) (model.Stack, error)
		CreateModuleVersion(module, source, version string) error
	}

	controllerImpl struct {
		log *zap.SugaredLogger

		stackRepo  stackRepository.Repository
		moduleRepo moduleRepository.Repository
	}
)
