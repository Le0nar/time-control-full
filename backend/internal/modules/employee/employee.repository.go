package employee

import (
	"github.com/jmoiron/sqlx"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

// TODO: mb move table name somewhere
const employeesTable = "employees"

func (r *EmployeeRepository) CreateEmployee(createEmployeeDto CreateEmployeeDto) (Employee, error) {
	var employee Employee
	
 	// TODO: set new employee to DB

	// query := fmt.Sprintf(
	// 	"INSERT INTO %s (email, name, password_hash) values ($1, $2, $3) RETURNING  id, email, name",
	// 	employeesTable,
	// )

	// row := r.db.QueryRow(query, createEmployeeDto.Email, createEmployeeDto.Name, createEmployeeDto.Password)
	// if err := row.Scan(&employee.Id, &employee.Email, &employee.Name); err != nil {
	// 	return employee, err
	// }

	return employee, nil
}