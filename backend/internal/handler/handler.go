package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control/internal/modules/company"
	"github.com/le0nar/time-control/internal/service"
)

type Handler struct {
	// TODO: mb use private literation for "CompanyHandler"
	CompanyHandler company.CompanyHandler
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		CompanyHandler: *company.NewCompanyHandler(services.CompanyService),
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	
	auth := router.Group("/auth")  
	{
		auth.POST("/company/sign-up", h.CompanyHandler.SignUp)
		auth.POST("/company/sign-in", h.CompanyHandler.SignIn )

		// auth.POST("/employee/sign-up")
		// auth.POST("/employee/sign-in")
	}

	// TODO: add api/company and api/employee endpointss

	return router
}
