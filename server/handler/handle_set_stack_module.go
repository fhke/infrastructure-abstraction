package handler

import (
	"net/http"

	"github.com/fhke/infrastructure-abstraction/server/handler/model/request"
	"github.com/fhke/infrastructure-abstraction/server/handler/model/response"
	"github.com/fhke/infrastructure-abstraction/server/storage/stack/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/samber/lo"
)

func (h *Handler) HandleSetStackModules(ctx *gin.Context) {
	log := h.log.With("request", "set_stack_modules")

	req := new(request.PatchStack)
	if err := ctx.ShouldBindWith(req, binding.JSON); err != nil {
		log.Warnw("Error binding request", "error", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	st, err := h.ctrl.SetStackModules(ctx, req.Name, req.Repository, req.ModuleVersions)
	if err != nil {
		log.Errorw("Error updating stack", "error", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, response.Stack{
		Modules: lo.MapValues(st.Modules, func(mod model.Module, _ string) response.StackModule {
			return response.StackModule{
				Version: mod.Version,
				Source:  mod.Source,
			}
		}),
	})
}
