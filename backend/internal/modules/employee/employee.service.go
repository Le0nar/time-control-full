package employee

import (
	"errors"
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

func (s *EmployeeService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &employeeTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		signingKey := os.Getenv("SIGNING_KEY")

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*employeeTokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *employeeTokenClaims")
	}

	return claims.EmployeeId, nil
}
