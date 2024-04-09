package handler

import (
	"fmt"
	"net/http"

	"github.com/fhke/infrastructure-abstraction/server/controller"
	"github.com/fhke/infrastructure-abstraction/server/handler/model/request"
	"github.com/fhke/infrastructure-abstraction/server/handler/model/response"
	"github.com/fhke/infrastructure-abstraction/server/storage/stack/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/samber/lo"
	"k8s.io/apimachinery/pkg/util/sets"
)

func (h *Handler) HandleBuildStack(ctx *gin.Context) {
	log := h.log.With("request", "build_stack")

	// parse & validate request
	req := new(request.BuildStack)
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
	stack, err := h.ctrl.BuildStack(req.Name, req.Repository, sets.New(req.ModuleNames...))
	if err != nil {
		if _, ok := err.(controller.ErrStackNotFound); ok {
			ctx.JSON(http.StatusNotFound, response.Error{Message: "stack not found"})
			return
		}
		if err, ok := err.(controller.ErrModuleNotFound); ok {
			ctx.JSON(http.StatusBadRequest, response.Error{Message: fmt.Sprintf("module %q not found", err.Name)})
			return
		}
		// unknown error
		log.Errorw("Error building stack", "error", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	// send response
	ctx.JSON(http.StatusOK, response.BuildStack{
		Modules: lo.MapValues(stack.Modules, func(mod model.Module, _ string) response.BuildStackModule {
			return response.BuildStackModule{
				Source:  mod.Source,
				Version: mod.Version,
			}
		}),
	})
}
