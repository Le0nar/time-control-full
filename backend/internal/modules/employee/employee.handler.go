package employee

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control/util"
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
	
	token, err := h.employeeService.GenerateEmployeeToken(signInEmployeeDto.Email, signInEmployeeDto.Password)
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
	if header == "" {
		util.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		util.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	employeeId, err := h.employeeService.ParseToken(headerParts[1]) 
	if err != nil {
		util.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	
	c.Set(employeeCtx, employeeId)
}
