package company

import (
	"errors"

	"github.com/le0nar/time-control/util"
)

type CompanyService struct {
	companyRepository CompanyRepository
}

func NewCompanyService(companyRepository CompanyRepository) *CompanyService {
	return &CompanyService{companyRepository: companyRepository}
}

const companyRole = "company"

func (s *CompanyService) CreateCompany(createCompanyDto CreateCompanyDto) (Company, error) {
	createCompanyDto.Password = util.GeneratePasswordHash(createCompanyDto.Password)
	return s.companyRepository.CreateCompany(createCompanyDto)
}

func (s *CompanyService) GetToken(email, password string) (string, error) {
	employee, err := s.companyRepository.GetCompany(email, util.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	return util.GenerateToken(employee.Id, companyRole)
}

func (s *CompanyService) GetCompanyId(token string) (int, error) {
	id, role, err := util.ParseToken(token)

	if role != companyRole {
		return 0, errors.New("access for this role is denied")
	}

	return id, err
}
