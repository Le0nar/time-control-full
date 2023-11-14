package activity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ActivityRepository struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

// TODO: rename "year" to "year_activity" (for "month" and "day" do the same)
// TODO: add references to db values
const (
	activityTable = "activity"
	yearTable = "year"
	monthTable = "month"
	dayTable = "day"
)

func (ar *ActivityRepository) CreateActivity(employeeId int, wasActive bool, checkDuration int64) (Activity, error) {
	var activity Activity

	query := fmt.Sprintf(
		"INSERT INTO %s (id, was_active, check_duration, employee_id, check_time) values ($1, $2, $3, $4, $5) RETURNING  id, was_active, employee_id",
		activityTable,
	)

	row := ar.db.QueryRow(query, uuid.New(), wasActive, checkDuration, employeeId, time.Now())

	err := row.Scan(&activity.Id, &activity.WasActive, &activity.EmployeeId)
	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (ar *ActivityRepository) ConfirmActivity(id string, checkDuration int64) error {
	query := fmt.Sprintf("UPDATE %s SET was_active='t' WHERE id='%s'", activityTable, id)

	_, err := ar.db.Exec(query)

	return err
}
