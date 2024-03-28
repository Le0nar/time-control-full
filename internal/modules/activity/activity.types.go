package activity

type DayActivity struct {
	Id           int    `json:"id" binding:"required" db:"id"`
	EmployeeId   int    `json:"employeeId" binding:"required" db:"employee_id"`
	ActivityDate string `json:"activityDate" binding:"required" db:"activity_date"`
	ActivityTime int64  `json:"activityTime" binding:"required" db:"activity_time"`
}

type DayActivityDto struct {
	Day          int   `json:"day"`
	ActivityTime int64 `json:"activityTime" db:"activity_time"`
}

type AddingWorkingTimeDto struct {
	EmployeeId   int    `json:"employeeId" binding:"required" db:"employee_id"`
	ActivityDate string `json:"activityDate" binding:"required" db:"activity_date"`
	ActivityTime int64  `json:"activityTime" binding:"required" db:"activity_time" `
}
