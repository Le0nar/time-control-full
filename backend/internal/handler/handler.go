package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control/internal/service"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	
	auth := router.Group("/auth")  
	{
		auth.POST("/company/sign-in")
		// auth.POST("/employee/sign-in")

		auth.POST("/company/sign-up")
		// auth.POST("/employee/sign-up")
	}

	// TODO: add api/company and api/employee endpointss

	return router
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
