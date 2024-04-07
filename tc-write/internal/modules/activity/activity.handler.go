package activity

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (ah *ActivityHandler) CheckActivity(c *gin.Context) {
	var checkingActivityDto CheckingActivityDto

	if err := c.BindJSON(&checkingActivityDto); err != nil {
		fmt.Printf("checkingActivityDto: %v\n", checkingActivityDto)
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: add check is photo exist

	activityEvent, err := ah.activityService.CreateCheckingActivityEvent(checkingActivityDto)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Save active time in read service (implementation of Eventual Consistancy)
	if activityEvent.WasActive {
		syncWithReadSerivce(c, activityEvent)
	}

	c.JSON(http.StatusOK, activityEvent)
}

func syncWithReadSerivce (c *gin.Context, activityEvent ActivityEvent) {
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

func (ah *ActivityHandler) ConfirmActivity(c *gin.Context) {
	var dto ConfirmingActivityDto
	if err := c.BindJSON(&dto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	activityEvent, err := ah.activityService.CreateConfirmingActivityEvent(dto)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	syncWithReadSerivce(c, activityEvent)

	c.JSON(http.StatusOK, "resource created or updated successfully")
}
