package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control-read/internal/modules/activity"
	"github.com/le0nar/time-control-read/internal/modules/gateway"
	"github.com/le0nar/time-control-read/internal/service"
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
	
	readService := router.Group("/read-service")
	{
		employeeGroup := readService.Group("/employee", h.GatewayHandler.IdentityEmployee)
		{
			employeeGroup.GET("/:id/activity", h.ActivityHandler.GetEmployeeMonthActivity)
			employeeGroup.POST("/activity", h.ActivityHandler.AddWorkTime)
		}
		
		// TODO: add group for company (get moonthly activity of employees)
	}
	
	return router
}
