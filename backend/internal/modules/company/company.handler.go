package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control/internal/util"
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
		return
	}
	
	token, err := h.companyService.GetToken(signInCompanyDto.Email, signInCompanyDto.Password)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

const (
	authorizationHeader = "Authorization"
	companyCtx = "companyId"
)

func (h *CompanyHandler) IdentityCompany(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	token, err := util.GetTokenFromHeader(header)
	if err != nil {
		util.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	companyId, err := h.companyService.GetCompanyId(token) 
	if err != nil {
		util.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	
	c.Set(companyCtx, companyId)
}
