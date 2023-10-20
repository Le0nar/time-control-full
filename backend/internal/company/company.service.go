package company

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CompanyService struct {
	companyRepository CompanyRepository
}

func NewCompanyService(companyRepository CompanyRepository) *CompanyService {
	return &CompanyService{companyRepository: companyRepository}
}

func (s *CompanyService) CreateCompany(createCompanyDto CreateCompanyDto) (Company, error) {
	createCompanyDto.Password = generatePasswordHash(createCompanyDto.Password)
	return s.companyRepository.CreateCompany(createCompanyDto)
}

type companyTokenClaims struct {
	jwt.StandardClaims
	CompanyId int `json:"user_id"`
}

// TODO: mb move to constants file
const tokenTTL = 12 * time.Hour

func (s *CompanyService) GenerateCompanyToken(email, password string) (string, error) {
	company, err := s.companyRepository.GetCompany(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &companyTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		company.Id,
	})
	signingKey := os.Getenv("SIGNING_KEY")

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	salt := os.Getenv("SALT")

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}