package company

import (
	"net/http"
	"strings"

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
	if header == "" {
		util.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		util.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	companyId, err := h.companyService.GetCompanyId(headerParts[1]) 
	if err != nil {
		util.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	
	c.Set(companyCtx, companyId)
}
