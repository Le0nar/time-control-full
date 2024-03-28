package activity

import (
	"os"
	"time"

	"github.com/google/uuid"
)

type ActivityEvent struct {
	Id    uuid.UUID    `json:"id" db:"id"`
	EmployeeId int `json:"employeeId" db:"employee_id"`
	CheckDuration int64 `json:"checkDuration" db:"check_duration"`
	CheckTime time.Time `json:"checkedTime" db:"checked_time"`
	WasActive bool  `json:"wasActive" db:"was_active"`
	EventTypeId int `json:"eventTypeId" db:"event_type_id"`
}

// if necessary check, can find couple checking event and confirm event by CheckTime and EmployeeId

type ActivityEventDto struct {
	EmployeeId int `json:"employeeId" binding:"required" db:"employee_id"`
	CheckDuration int64 `json:"checkDuration" binding:"required" db:"check_duration"`
	CheckTime time.Time `json:"checkedTime" binding:"required" db:"checked_time"`
	WasActive bool `json:"wasActive" binding:"required" db:"was_active"`
	EventTypeId int `json:"eventTypeId" binding:"required" db:"event_type_id"`
}

type CheckingActivityDto struct {
	Photo os.File `json:"photo" binding:"required"`
	InactivityTime int64 `json:"inactivityTime" binding:"required"`
	EmployeeId int `json:"employeeId" binding:"required"`
}

type ConfirmingActivityDto struct {
	EmployeeId int `json:"employeeId" binding:"required"`
	CheckDuration int64 `json:"checkDuration" binding:"required"`
	CheckTime time.Time `json:"checkedTime" binding:"required"`
}

type AddingWorkingTimeDto struct {
	EmployeeId   int    `json:"employeeId" binding:"required" db:"employee_id"`
	ActivityDate string `json:"activityDate" binding:"required" db:"activity_date"`
	ActivityTime int64  `json:"activityTime" binding:"required" db:"activity_time" `
}
