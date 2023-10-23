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

const companyTable = "company"

func (r *CompanyRepository) CreateCompany(createCompanyDto CreateCompanyDto) (Company, error) {
	var company Company
	
	query := fmt.Sprintf("INSERT INTO %s (email, name, password_hash) values ($1, $2, $3) RETURNING  id, email, name", companyTable)

	row := r.db.QueryRow(query, createCompanyDto.Email, createCompanyDto.Name, createCompanyDto.Password)
	if err := row.Scan(&company.Id, &company.Email, &company.Name); err != nil {
		return company, err
	}

	return company, nil
}

func (r *CompanyRepository) GetCompany(email, passwordHash string) (Company, error) {
	var company Company

	query := fmt.Sprintf("SELECT id, email, name FROM %s WHERE email=$1 and password_hash=$2", companyTable)
	err := r.db.Get(&company, query, email, passwordHash)

	return company, err
}
