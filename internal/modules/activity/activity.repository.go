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

func (ar *ActivityRepository) CreateActivityEvent(activityEventDto ActivityEventDto) (ActivityEvent, error) {
	var activityEvent ActivityEvent
	
	query := fmt.Sprintf(
		"INSERT INTO %s (employee_id, check_duration, check_time, was_active, event_type_id) values ($1, $2, $3, $4, $5) RETURNING *",
		activityEventTablet,
	)

	row := ar.db.QueryRow(
		query,
		activityEventDto.EmployeeId,
		activityEventDto.CheckDuration,
		activityEventDto.CheckTime,
		activityEventDto.WasActive,
		activityEventDto.EventTypeId,
	)

	err := row.Scan(
		&activityEvent.Id,
		&activityEvent.EmployeeId,
		&activityEvent.CheckDuration,
		&activityEvent.CheckTime,
		&activityEvent.WasActive,
		&activityEvent.EventTypeId,
	)

	return activityEvent, err
}
