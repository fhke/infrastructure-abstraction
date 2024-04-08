package handler

import (
	"github.com/fhke/infrastructure-abstraction/server/controller"
	"go.uber.org/zap"
)

type Handler struct {
	log  *zap.SugaredLogger
	ctrl controller.Controller
}

func New(log *zap.SugaredLogger, ctrl controller.Controller) *Handler {
	return &Handler{
		log:  log,
		ctrl: ctrl,
	}
}
