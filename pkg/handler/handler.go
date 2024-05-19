package handler

import (
	"ecurrency/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoute() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/rate", h.getRates)
		api.POST("/subscription", h.getEmail)
	}

	return router
}
