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

type YearActivity struct {
	Id    int    `json:"id" db:"id" `
	EmployeeId int `json:"employeeId" db:"employee_id"`
	Year    int    `json:"year" db:"year" `
}

type MonthActivity struct {
	Id    int    `json:"id" db:"id" `
	YearId int `json:"yearId" db:"year_id"`
	Month    int    `json:"month" db:"month"`
}

type DayActivity struct {
	// TODO: mb change type of Id to int64 or uuid
	Id    int    `json:"id" db:"id" `
	MonthId int `json:"monthId" db:"month_id"`
	Day    int    `json:"day" db:"day"`
	ActivityTime int64  `json:"activityTime" db:"activity_time"`
}
