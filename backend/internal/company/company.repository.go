package company

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CompanyRepository struct {
	db *sqlx.DB
}

func NewCompanyRepository(db *sqlx.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

// TODO: mb move table name somewhere
const companiesTable = "companies"

func (r *CompanyRepository) CreateCompany(createCompanyDto CreateCompanyDto) (CompanyDto, error) {
	var company CompanyDto
	
	query := fmt.Sprintf("INSERT INTO %s (email, name, password_hash) values ($1, $2, $3) RETURNING  id, email, name", companiesTable)

	row := r.db.QueryRow(query, createCompanyDto.Email, createCompanyDto.Name, createCompanyDto.Password)
	if err := row.Scan(&company.Id, &company.Email, &company.Name); err != nil {
		return company, err
	}

	return company, nil
}
