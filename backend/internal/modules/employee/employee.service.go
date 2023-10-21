package employee

import (
	"errors"

	"github.com/le0nar/time-control/internal/util"
)

type EmployeeService struct {
	employeeRepository EmployeeRepository
}

func NewEmployeeService(employeeRepository EmployeeRepository) *EmployeeService {
	return &EmployeeService{employeeRepository: employeeRepository}
}

const employeeRole = "employee"

func (s *EmployeeService) CreateEmployee(createEmployeeDto CreateEmployeeDto) (Employee, error) {
	createEmployeeDto.Password = util.GeneratePasswordHash(createEmployeeDto.Password)
	return s.employeeRepository.CreateEmployee(createEmployeeDto)
}

func (s *EmployeeService) GetToken(email, password string) (string, error) {
	employee, err := s.employeeRepository.GetEmployee(email, util.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	return util.GenerateToken(employee.Id, employeeRole)
}

func (s *EmployeeService) GetEmployeeId(token string) (int, error) {
	id, role, err := util.ParseToken(token)

	if role != employeeRole {
		return 0, errors.New("access for this role is denied")
	}

	return id, err
}
