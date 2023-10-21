package activity

import (
	"os"

	"github.com/google/uuid"
)

// TODO: add json
type Activity struct {
	Id    uuid.UUID    `json:"id" db:"id"`
	WasActive bool  `json:"wasActive" db:"was_active"`
	EmployeeId int `json:"employeeId" db:"employee_id"`
	// CheckedTime int64 `db:"checked_time"`
}

type CreateActivityDto struct {
	Photo os.File `json:"photo"`
	InactivityTime int64 `json:"inactivityTime"`
	EmployeeId int `json:"employeeId"`
}
