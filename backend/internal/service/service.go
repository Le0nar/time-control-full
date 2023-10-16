package service

import (
	"github.com/le0nar/time-control/internal/company"
	"github.com/le0nar/time-control/internal/repository"
)

type Service struct {
	CompanyService company.CompanyService
}

func NewService(repository *repository.Repository) *Service {

	return &Service{
		CompanyService: *company.NewCompanyService(repository.CompanyRepository),
	}
}
