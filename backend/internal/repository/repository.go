package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/le0nar/time-control/internal/modules/activity"
	"github.com/le0nar/time-control/internal/modules/company"
	"github.com/le0nar/time-control/internal/modules/employee"
)

type Repository struct {
	CompanyRepository company.CompanyRepository
	EmployeeRepository employee.EmployeeRepository
	ActivityRepository activity.ActivityRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CompanyRepository: *company.NewCompanyRepository(db),
		EmployeeRepository: *employee.NewEmployeeRepository(db),
		ActivityRepository: *activity.NewActivityRepository(db),
	}
}
