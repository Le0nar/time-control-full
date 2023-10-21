package service

import (
	"github.com/le0nar/time-control/internal/modules/activity"
	"github.com/le0nar/time-control/internal/modules/company"
	"github.com/le0nar/time-control/internal/modules/employee"
	"github.com/le0nar/time-control/internal/repository"
)

type Service struct {
	CompanyService company.CompanyService
	EmployeeService employee.EmployeeService
	ActivitySerice activity.ActivityService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		CompanyService: *company.NewCompanyService(repository.CompanyRepository),
		EmployeeService: *employee.NewEmployeeService(repository.EmployeeRepository),
		ActivitySerice: *activity.NewActiviySerivce(repository.ActivityRepository),
	}
}
