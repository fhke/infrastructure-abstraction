package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Register(routes gin.IRoutes) {
	routes.POST("/api/stack/build", h.HandleBuildStack)
	routes.POST("/api/modules", h.HandleCreateModuleVersion)
	routes.GET("/api/modules", h.HandleGetModuleNames)
	routes.PATCH("/api/stack", h.HandleSetStackModules)
}
