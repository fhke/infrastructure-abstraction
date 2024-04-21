package handler

import (
	"net/http"

	"github.com/fhke/infrastructure-abstraction/server/handler/model/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleGetModuleNames(ctx *gin.Context) {
	log := h.log.With("request", "get_module_names")

	names, err := h.ctrl.GetModuleNames(ctx)
	if err != nil {
		log.Errorw("Error getting module names", "error", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, response.ModuleNames{
		Names: names,
	})
}
