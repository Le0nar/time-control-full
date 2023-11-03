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

	company := auth.Group("/company")
	{
		company.POST("/sign-up", h.CompanyHandler.SignUp)
		company.POST("/sign-in", h.CompanyHandler.SignIn)
		// company.POST("/validate", h.CompanyHandler.IdentityCompany)
	}

	employee := auth.Group("/employee")
	{
		employee.POST("/sign-up", h.EmployeeHandler.SignUp)
		employee.POST("/sign-in", h.EmployeeHandler.SignIn)
		// TODO: mb use .GET
		employee.POST("/validate", h.EmployeeHandler.IdentityEmployee)
	}

	return router
}
