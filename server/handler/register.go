package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Register(e *gin.Engine) {
	e.POST("/api/stack/build", h.HandleBuildStack)
	e.POST("/api/modules", h.HandleCreateModuleVersion)
	e.GET("/api/modules", h.HandleGetModuleNames)
	e.PATCH("/api/stack", h.HandleSetStackModules)
}
