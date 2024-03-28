package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control/internal/modules/activity"
	"github.com/le0nar/time-control/internal/modules/gateway"
	"github.com/le0nar/time-control/internal/service"
)

type Handler struct {
	ActivityHandler activity.ActivityHandler
	GatewayHandler gateway.GatewayHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		ActivityHandler: *activity.NewActivityHandler(service.ActivitySerice),
		GatewayHandler: *gateway.NewGatewayHandler(),
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	
	writeService := router.Group("/write-service", h.GatewayHandler.IdentityEmployee)  
	{
		writeService.POST("/check-activity", h.ActivityHandler.CheckActivity)
		writeService.POST("/confirm-activity/", h.ActivityHandler.ConfirmActivity)
	}

	return router
}
