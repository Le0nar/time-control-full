package company

type Company struct {
	// TODO: add password here
	Id    int    `json:"id" db:"id"`
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type CreateCompanyDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type SignInCompanyDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
