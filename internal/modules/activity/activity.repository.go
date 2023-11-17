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
	activityEventTablet = "activity_event"
)

func (ar *ActivityRepository) CreateActivityEvent(activityEvent ActivityEvent) error {
// 2) в хендлере вызвать метод add work from read service
// 3) create another endpoint for write-serivce without checking (when user clicked "i'm here")

	query := fmt.Sprintf(
		"INSERT INTO %s (employee_id, check_duration, check_time, was_active, event_type_id) values ($1, $2, $3, $4, $5) RETURNING was_active",
		activityEventTablet,
	)

	row := ar.db.QueryRow(
		query,
		activityEvent.EmployeeId,
		activityEvent.CheckDuration,
		activityEvent.CheckTime,
		activityEvent.WasActive,
		activityEvent.EventTypeId,
	)

	err := row.Err()

	return err
}
