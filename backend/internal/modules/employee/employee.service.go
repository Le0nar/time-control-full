package employee

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/le0nar/time-control/util"
)

type EmployeeService struct {
	employeeRepository EmployeeRepository
}

func NewEmployeeService(employeeRepository EmployeeRepository) *EmployeeService {
	return &EmployeeService{employeeRepository: employeeRepository}
}

func (s *EmployeeService) CreateEmployee(createEmployeeDto CreateEmployeeDto) (Employee, error) {
	createEmployeeDto.Password = util.GeneratePasswordHash(createEmployeeDto.Password)
	return s.employeeRepository.CreateEmployee(createEmployeeDto)
}

type employeeTokenClaims struct {
	jwt.StandardClaims
	EmployeeId int `json:"employee_id"`
}

// TODO: mb move to constants file
const tokenTTL = 12 * time.Hour

func (s *EmployeeService) GenerateEmployeeToken(email, password string, ) (string, error) {
	employee, err := s.employeeRepository.GetEmployee(email, util.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &employeeTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		employee.Id,
	})
	signingKey := os.Getenv("SIGNING_KEY")

	return token.SignedString([]byte(signingKey))
}
