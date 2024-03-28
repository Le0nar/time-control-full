package service

import (
	"time-control-auth/internal/modules/company"
	"time-control-auth/internal/modules/employee"
	"time-control-auth/internal/repository"
)

type Service struct {
	CompanyService company.CompanyService
	EmployeeService employee.EmployeeService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		CompanyService: *company.NewCompanyService(repository.CompanyRepository),
		EmployeeService: *employee.NewEmployeeService(repository.EmployeeRepository),
	}
}
