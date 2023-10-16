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
	var dto CreateCompanyDto

	if err := c.BindJSON(&dto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	company, err := h.companyService.CreateCompany(dto)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, company)
}
