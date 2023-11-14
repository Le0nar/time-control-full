package activity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/le0nar/time-control/internal/util"
)

type ActivityHandler struct {
	activityService ActivityService
}

func NewActivityHandler(activityService ActivityService) *ActivityHandler {
	return &ActivityHandler{activityService: activityService}
}

// TODO: rename
func (ah *ActivityHandler) CreateActivity(c *gin.Context) {
	var createActivityDto CreateActivityDto

	if err := c.BindJSON(&createActivityDto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	activity, err := ah.activityService.CreateActivity(createActivityDto)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Send activity to client if user was unactive
	if !activity.WasActive {
		c.JSON(http.StatusOK, activity)
		return
	}

	// TODO: call http method from read service
	// Eventual Consistancy
	// err = ah.activityService.AddWorkTime(activity.EmployeeId)
	// if err != nil {
	// 	util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusCreated, "resource created successfully")
}

func (ah *ActivityHandler) ConfirmActivity(c *gin.Context) {
	id := c.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	// TODO: create new event instead of update
	err = ah.activityService.ConfirmActivity((id))
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: move to read-service
	// TODO: get employeeId from request
	// err = ah.activityService.AddWorkTime(activity.EmployeeId)
	// if err != nil {
	// 	util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusOK, "resource updated successfully")
}
