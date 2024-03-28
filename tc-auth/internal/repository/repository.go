package repository

import (
	"fmt"

	"time-control-auth/internal/modules/company"
	"time-control-auth/internal/modules/employee"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	CompanyRepository company.CompanyRepository
	EmployeeRepository employee.EmployeeRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CompanyRepository: *company.NewCompanyRepository(db),
		EmployeeRepository: *employee.NewEmployeeRepository(db),
	}
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// TODO: mb move function to another file within the "repository" module
func NewPostgresDB(cfg DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", 
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
