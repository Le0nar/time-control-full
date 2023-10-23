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

const activityTable = "activity"

// TODO: mb return *Activity
func (r *ActivityRepository) CreateActivity(employeeId int, wasActive bool, checkDuration int64) (Activity, error) {
	var activity Activity

	query := fmt.Sprintf(
		"INSERT INTO %s (id, was_active, check_duration, employee_id, check_time) values ($1, $2, $3, $4, $5) RETURNING  id, was_active, employee_id",
		activityTable,
	)

	// TODO: mb create struct for Activity in Database
	// TOOD: mb move time.Now & uuid.New to serivce (but for what?)
	row := r.db.QueryRow(query, uuid.New(), wasActive, checkDuration, employeeId, time.Now())

	err := row.Scan(&activity.Id, &activity.WasActive, &activity.EmployeeId)
	if err != nil {
		return activity, err
	}

	return activity, nil
}
