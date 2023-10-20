package employee

import (
	"net/http"

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
