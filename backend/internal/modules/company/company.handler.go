package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control/util"
)

type CompanyHandler struct {
	companyService CompanyService
}

func NewCompanyHandler(companyService CompanyService) *CompanyHandler {
	return &CompanyHandler{companyService: companyService}
}

func (h *CompanyHandler) SignUp(c *gin.Context) {
	var createCompanyDto CreateCompanyDto

	if err := c.BindJSON(&createCompanyDto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	company, err := h.companyService.CreateCompany(createCompanyDto)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) SignIn(c *gin.Context) {
	var signInCompanyDto SignInCompanyDto

	if err := c.BindJSON(&signInCompanyDto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	
	token, err := h.companyService.GenerateCompanyToken(signInCompanyDto.Email, signInCompanyDto.Password)
		if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
