package activity

import (
	"os"

	"github.com/google/uuid"
)

type Activity struct {
	Id    uuid.UUID    `json:"id" db:"id" `
	WasActive bool  `json:"wasActive" db:"was_active"`
	EmployeeId int `json:"employeeId" db:"employee_id"`
	// CheckedTime int64 `json:"checkedTime" db:"checked_time"`
}

type CreateActivityDto struct {
	Photo os.File `json:"photo" binding:"required"`
	InactivityTime int64 `json:"inactivityTime" binding:"required"`
	EmployeeId int `json:"employeeId" binding:"required"`
}
