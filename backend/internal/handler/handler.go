package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control/internal/modules/activity"
	"github.com/le0nar/time-control/internal/modules/company"
	"github.com/le0nar/time-control/internal/modules/employee"
	"github.com/le0nar/time-control/internal/service"
)

type Handler struct {
	// TODO: mb use private literation for handlers of modules
	CompanyHandler company.CompanyHandler
	EmployeeHandler employee.EmployeeHandler
	ActivityHandler activity.ActivityHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		CompanyHandler: *company.NewCompanyHandler(service.CompanyService),
		EmployeeHandler: *employee.NewEmployeeHandler(service.EmployeeService),
		ActivityHandler: *activity.NewActivityHandler(service.ActivitySerice),
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	
	// TODO: mb wrap all to "/api"

	auth := router.Group("/auth")  
	{
		auth.POST("/company/sign-up", h.CompanyHandler.SignUp)
		auth.POST("/company/sign-in", h.CompanyHandler.SignIn)

		auth.POST("/employee/sign-up", h.EmployeeHandler.SignUp)
		auth.POST("/employee/sign-in", h.EmployeeHandler.SignIn)
	}

	employeeApi := router.Group("/employee-api")  
	{
		employeeApi.POST("/activity", h.ActivityHandler.CreateActivity)
		// TODO: make activity with id :id wasActive: true
		// employeeApi.PATCH("/activity/:id", h.updateNews)
	}

	return router
}
