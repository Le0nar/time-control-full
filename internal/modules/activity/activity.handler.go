package activity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control-read/internal/util"
)

type ActivityHandler struct {
	activityService ActivityService
}

func NewActivityHandler(activityService ActivityService) *ActivityHandler {
	return &ActivityHandler{activityService: activityService}
}

func (ah *ActivityHandler) GetEmployeeMonthActivity(c *gin.Context) {
	employeeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	yearQuery, ok := c.GetQuery("year")
	if !ok {
		util.NewErrorResponse(c, http.StatusBadRequest, "missing year query param")
		return
	}

	year, err := strconv.Atoi(yearQuery)
	if err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, "invalid year query param")
		return
	}

	monthQuery, ok := c.GetQuery("month")
	if !ok {
		util.NewErrorResponse(c, http.StatusBadRequest, "missing month query param")
		return
	}

	month, err := strconv.Atoi(monthQuery)
	if err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, "invalid month query param")
		return
	}

	monthActivityDto, err := ah.activityService.GetEmployeeMonthActivity(employeeId, year, month)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, monthActivityDto)
}


// TODO: add work time endpoint