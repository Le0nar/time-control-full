package employee

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}



const employeesTable = "employees"

func (r *EmployeeRepository) CreateEmployee(createEmployeeDto CreateEmployeeDto) (Employee, error) {
	var employee Employee

	const insertedFields  = "email, password_hash, first_name, second_name, patronymic, company_id"
	const returnedFields = "email, id, first_name, second_name, patronymic, company_id"

	query := fmt.Sprintf(
		"INSERT INTO %s (%s)  values ($1, $2, $3, $4, $5, $6) RETURNING  %s",
		employeesTable,
		insertedFields,
		returnedFields,
	)

	row := r.db.QueryRow(
		query,
		createEmployeeDto.Email,
		createEmployeeDto.Password,
		createEmployeeDto.FirstName,
		createEmployeeDto.SecondName,
		createEmployeeDto.Patronymic,
		createEmployeeDto.CompanyId,
	)

	err := row.Scan(
		&employee.Email,
		&employee.Id,
		&employee.FirstName,
		&employee.SecondName,
		&employee.Patronymic,
		&employee.CompanyId,
	)
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *EmployeeRepository) GetEmployee(email, passwordHash string) (Employee, error) {
	var employee Employee

	query := fmt.Sprintf(
		"SELECT id, email, first_name, first_name, patronymic, company_id FROM %s WHERE email=$1 and password_hash=$2",
	 	employeesTable,
	)
	err := r.db.Get(&employee, query, email, passwordHash)

	return employee, err
}
