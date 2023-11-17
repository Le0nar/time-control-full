package activity

import (
	"bytes"
	"encoding/json"
	"io"
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

const (
	authorizationHeader = "Authorization"
	addingWorkTimeUrl = "http://localhost:8002/read-service/employee/activity"
	YYYYMMDD = "2006-01-02"
)

// TODO: rename
func (ah *ActivityHandler) CheckActivity(c *gin.Context) {
	var checkingActivityDto CheckingActivityDto

	if err := c.BindJSON(&checkingActivityDto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: add check is photo exist

	activityEvent, err := ah.activityService.CreateCheckingActivityEvent(checkingActivityDto)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: mb move to function and reuse in antoher event
	// Save active time in read service (implementation of Eventual Consistancy)
	if activityEvent.WasActive {
		var addingWorkingTimeDto AddingWorkingTimeDto

		// change format from time.Time to yyyy-mm-dd 
		addingWorkingTimeDto.ActivityDate = activityEvent.CheckTime.Format(YYYYMMDD)
		addingWorkingTimeDto.EmployeeId = activityEvent.EmployeeId
		addingWorkingTimeDto.ActivityTime = activityEvent.CheckDuration


		jsonStr, err := json.Marshal(addingWorkingTimeDto)
		if err != nil {
			util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		req, err := http.NewRequest(http.MethodPost, addingWorkTimeUrl, bytes.NewBuffer(jsonStr))
		if err != nil {
			util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		
		header := c.GetHeader(authorizationHeader)
		tokenValue, err := util.GetTokenFromHeader(header)
		if err != nil {
			util.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		token := "Bearer " + tokenValue

		req.Header.Set(authorizationHeader, token)

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			util.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
	
		body, _ := io.ReadAll(res.Body)
	
		if res.StatusCode != http.StatusOK {
			util.NewErrorResponse(c, http.StatusUnauthorized, string(body))
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"wasActive": activityEvent.WasActive,
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
