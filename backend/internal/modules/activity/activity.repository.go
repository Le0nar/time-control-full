package activity

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type ActivityRepository struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

const activitiesTable = "activities"

// TODO: mb return *Activity
func (r *ActivityRepository) CreateActivity(employeeId int, wasActive bool, checkDuration int64) (Activity, error) {
	var activity Activity

	query := fmt.Sprintf(
		"INSERT INTO %s (was_active, check_duration, employee_id, check_time) values ($1, $2, $3, $4) RETURNING  id, was_active, employee_id",
		activitiesTable,
	)

	// TOOD: mb move time.Now to serivce (but for what?)
	row := r.db.QueryRow(query, wasActive, checkDuration, employeeId, time.Now())
		if err := row.Scan(&activity.Id, &activity.WasActive, &activity.EmployeeId); err != nil {
		return activity, err
	}

	return activity, nil
}
