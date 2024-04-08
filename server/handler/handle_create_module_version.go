package handler

import (
	"net/http"

	"github.com/fhke/infrastructure-abstraction/server/controller"
	"github.com/fhke/infrastructure-abstraction/server/handler/model/request"
	"github.com/fhke/infrastructure-abstraction/server/handler/model/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) HandleCreateModuleVersion(ctx *gin.Context) {
	log := h.log.With("request", "create_module_version")

	// parse & validate request
	req := new(request.CreateModuleVersion)
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		log.Warnw("Error binding request", "error", err)
		ctx.Status(http.StatusBadRequest)
		return
	} else if err := req.Validate(); err != nil {
		log.Warnw("Request validation failure", "error", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	// build stack
	if err := h.ctrl.CreateModuleVersion(req.Name, req.Source, req.Version); err != nil {
		if _, ok := err.(controller.ErrModuleVersionAlreadyExists); ok {
			log.Warnw("Request to create pre-existing module & version", "module", req.Name, "version", req.Version.String())
			ctx.JSON(http.StatusConflict, response.Error{Message: "version for module already exists"})
			return
		}
		// unknown error
		log.Errorw("Error creating version for module", "error", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	// send response
	ctx.Status(http.StatusCreated)
}
