package employee

import "github.com/le0nar/time-control/util"

type EmployeeService struct {
	employeeRepository EmployeeRepository
}

func (s *EmployeeService) CreateEmployee(createEmployeeDto CreateEmployeeDto) (Employee, error) {
	createEmployeeDto.Password = util.GeneratePasswordHash(createEmployeeDto.Password)
	return s.employeeRepository.CreateEmployee(createEmployeeDto)
}
