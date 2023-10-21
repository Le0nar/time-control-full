package employee

type Employee struct {
	Id         int    `json:"id" db:"id"`
	Email      string `json:"email" binding:"required"`
	FirstName  string `json:"firstName" db:"first_name" binding:"required"`
	SecondName string `json:"secondName" db:"second_name" binding:"required"`
	Patronymic string `json:"patronymic"`
	CompanyId  int    `json:"companyId" db:"company_id" binding:"required"`
}

type CreateEmployeeDto struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	FirstName  string `json:"firstName" db:"first_name" binding:"required"`
	SecondName string `json:"secondName" db:"second_name" binding:"required"`
	Patronymic string `json:"patronymic"`
	CompanyId  int    `json:"companyId" db:"company_id" binding:"required"`
}

type SignInEmployeeDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
