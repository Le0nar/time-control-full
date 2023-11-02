package employee

import (
	"net/http"

	"time-control-auth/internal/util"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeService EmployeeService
}

func NewEmployeeHandler(employeeService EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{employeeService: employeeService}
}

func (h *EmployeeHandler) SignUp(c *gin.Context) {
	var createEmployeeDto CreateEmployeeDto

	if err := c.BindJSON(&createEmployeeDto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	company, err := h.employeeService.CreateEmployee(createEmployeeDto)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *EmployeeHandler) SignIn(c *gin.Context) {
	var signInEmployeeDto SignInEmployeeDto

	if err := c.BindJSON(&signInEmployeeDto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	token, err := h.employeeService.GetToken(signInEmployeeDto.Email, signInEmployeeDto.Password)
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
	employeeCtx = "employeeId"
)

func (h *EmployeeHandler) IdentityEmployee(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	token, err := util.GetTokenFromHeader(header)
	if err != nil {
		util.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	employeeId, err := h.employeeService.GetEmployeeId(token) 
	if err != nil {
		util.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	
	c.Set(employeeCtx, employeeId)
}
