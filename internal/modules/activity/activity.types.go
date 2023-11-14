package activity

type DayActivity struct {
	Id           int    `json:"id" db:"id"`
	EmployeeId   int    `json:"employeeId" db:"employee_id"`
	ActivityDate string `json:"activityDate" db:"activity_date"`
	ActivityTime int64  `json:"activityTime" db:"activity_time"`
}

type DayActivityDto struct {
	Day          int   `json:"day"`
	ActivityTime int64 `json:"activityTime"`
}

type MonthActivityDto struct {
	Month int              `json:"month"`
	Days  []DayActivityDto `json:"days"`
}
