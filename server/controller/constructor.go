package controller

import (
	moduleRepository "github.com/fhke/infrastructure-abstraction/server/storage/module/repository"
	stackRepository "github.com/fhke/infrastructure-abstraction/server/storage/stack/repository"
	"go.uber.org/zap"
)

func New(log *zap.SugaredLogger, stackRepo stackRepository.Repository, moduleRepo moduleRepository.Repository) Controller {
	return &controllerImpl{
		log:        log,
		stackRepo:  stackRepo,
		moduleRepo: moduleRepo,
	}
}
