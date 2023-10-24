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

func (h *ActivityHandler) CreateActivity(c *gin.Context) {
	var createActivityDto CreateActivityDto

	if err := c.BindJSON(&createActivityDto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	activity, err := h.activityService.CreateActivity(createActivityDto)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if activity.WasActive {
		c.JSON(http.StatusCreated, "resource created successfully")
		return
	}

	c.JSON(http.StatusOK, activity)
}

func (h *ActivityHandler) ConfirmActivity(c *gin.Context) {
	id := c.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.activityService.ConfirmActivity((id))
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "resource updated successfully")
}
