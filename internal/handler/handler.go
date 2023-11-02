package handler

import (
	"time-control-auth/internal/modules/company"
	"time-control-auth/internal/modules/employee"
	"time-control-auth/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	CompanyHandler company.CompanyHandler
	EmployeeHandler employee.EmployeeHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		CompanyHandler: *company.NewCompanyHandler(service.CompanyService),
		EmployeeHandler: *employee.NewEmployeeHandler(service.EmployeeService),
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth-service")  
	{
		auth.POST("/company/sign-up", h.CompanyHandler.SignUp)
		auth.POST("/company/sign-in", h.CompanyHandler.SignIn)

		auth.POST("/employee/sign-up", h.EmployeeHandler.SignUp)
		auth.POST("/employee/sign-in", h.EmployeeHandler.SignIn)
	}

	return router
}
