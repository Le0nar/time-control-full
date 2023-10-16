package company

type CreateCompanyDto struct {
	Email       string       `json:"email" binding:"required"`
	Password    string       `json:"password" binding:"required"`
	Name        string       `json:"name" binding:"required"`
}

// TODO: use name "Company" instead of CompanyDto
type CompanyDto struct {
	Id          int          `json:"id" db:"id"`
	Email       string       `json:"email" binding:"required"`
	Name        string       `json:"name" binding:"required"`
}

type SignInCompanyDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
