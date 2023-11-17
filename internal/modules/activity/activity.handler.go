package activity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control/internal/util"
)

type ActivityHandler struct {
	activityService ActivityService
}

func NewActivityHandler(activityService ActivityService) *ActivityHandler {
	return &ActivityHandler{activityService: activityService}
}

// TODO: rename
func (ah *ActivityHandler) CheckActivity(c *gin.Context) {
	var checkingActivityDto CheckingActivityDto

	if err := c.BindJSON(&checkingActivityDto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: add check is photo exist

	wasEmployeeActive, err := ah.activityService.CreateActivityEvent(checkingActivityDto)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Save active time in read service
	// if wasEmployeeActive {
		// TODO: call http method from read service
		// Eventual Consistancy
		// err = ah.activityService.AddWorkTime(activityEvent.EmployeeId)
		// if err != nil {
		// 	util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		// 	return
		// }
	// }

	c.JSON(http.StatusOK, map[string]interface{}{
		"wasActive": wasEmployeeActive,
	})
}

// func (ah *ActivityHandler) ConfirmActivity(c *gin.Context) {
// 	id := c.Param("id")

// 	_, err := uuid.Parse(id)
// 	if err != nil {
// 		util.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
// 		return
// 	}

// 	// TODO: create new event instead of update
// 	err = ah.activityService.ConfirmActivity((id))
// 	if err != nil {
// 		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	// TODO: move to read-service
// 	// TODO: get employeeId from request
// 	// err = ah.activityService.AddWorkTime(activity.EmployeeId)
// 	// if err != nil {
// 	// 	util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 	// 	return
// 	// }

// 	c.JSON(http.StatusOK, "resource updated successfully")
// }
