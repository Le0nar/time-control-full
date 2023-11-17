package activity

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ActivityRepository struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

const (
	dayActivityTable = "day_activity"
	employeeIdColumn = "employee_id"
	activityDateColumn = "activity_date"
	activityTimeColumn = "activity_time"
)

func (ar *ActivityRepository) GetEmployeeMonthActivity(employeeId, year, month int) ([]DayActivityDto, error) {
	var dayActivityDtoList  []DayActivityDto

	query := fmt.Sprintf(
		"SELECT EXTRACT(DAY FROM %s) as day, %s from %s WHERE %s = $1 AND EXTRACT(MONTH FROM %s) = $2 AND EXTRACT(YEAR FROM %s) = $3 ORDER BY day;",
		activityDateColumn,
		activityTimeColumn,
		dayActivityTable,
		employeeIdColumn,
		activityDateColumn,
		activityDateColumn,
	)

	err := ar.db.Select(
		&dayActivityDtoList,
		query,
		employeeId,
		month,
		year,
	)

	return dayActivityDtoList, err
}

func (ar *ActivityRepository) AddWorkTime(dto AddingWorkingTimeDto) error {
	// Update day or create day if it doesnt exist

	var dayActivity DayActivity
	
	// Check is day activity exist
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE %s=$1 AND %s=$2",
		dayActivityTable,
		employeeIdColumn,
		activityDateColumn,
	)
	err := ar.db.Get(&dayActivity, query, dto.EmployeeId, dto.ActivityDate)

	// Create new day activity if it doesnt exist
	if err != nil {
		query := fmt.Sprintf(
			"INSERT INTO %s (employee_id, activity_date, activity_time) values ($1, $2, $3)",
			dayActivityTable,
		)
	
		row := ar.db.QueryRow(query, dto.EmployeeId, dto.ActivityDate, dto.ActivityTime)
		err := row.Err()

		return err
	}

	// Update day activity time if day activity already exists
	dayActivity.ActivityTime += dto.ActivityTime

	query = fmt.Sprintf("UPDATE %s SET activity_time=$1 WHERE id=$2", dayActivityTable)
	row := ar.db.QueryRow(
		query,
		dayActivity.ActivityTime,
		dayActivity.Id,
	)

	err = row.Err()

	return err
}
