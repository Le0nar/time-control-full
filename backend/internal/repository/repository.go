package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/le0nar/time-control/internal/modules/company"
)

type Repository struct {
	CompanyRepository company.CompanyRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CompanyRepository: *company.NewCompanyRepository(db),
	}
}
