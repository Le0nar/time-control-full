package company

import (
	"crypto/sha1"
	"fmt"
	"os"
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

// TODO: mb move function to company.utils.go
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	salt := os.Getenv("SALT")

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
