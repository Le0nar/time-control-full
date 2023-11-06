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

func (ar *ActivityRepository) CreateActivity(employeeId int, wasActive bool, checkDuration int64) (Activity, error) {
	var activity Activity

	query := fmt.Sprintf(
		"INSERT INTO %s (id, was_active, check_duration, employee_id, check_time) values ($1, $2, $3, $4, $5) RETURNING  id, was_active, employee_id",
		activityTable,
	)

	// TODO: mb create struct for Activity in Database
	// TOOD: mb move time.Now & uuid.New to serivce (but for what?)
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

func (ar * ActivityRepository) AddWorkTime(employeeId int) error {
	year, month, day := time.Now().Date()

	var yearActivity YearActivity


	
	// query := fmt.Sprintf("SELECT id, email, name FROM %s WHERE email=$1 and password_hash=$2", companyTable)
	// err := r.db.Get(&company, query, email, passwordHash)

	// return company, err

	// 1) get year, or create year if it doesnt exist
	// 2) get month or create month if it doesnt exist
	// 3) getday or create day if it doesnt exist
	// 4) update day activity time
	
	return nil
}