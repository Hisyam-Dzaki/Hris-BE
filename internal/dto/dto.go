package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
	CompanyID uint  `json:"company_id" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

type CreateCompanyRequest struct {
	Name    string `json:"name" binding:"required"`
	Slug    string `json:"slug" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Website string `json:"website"`
}

type UpdateCompanyRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Website string `json:"website"`
	Logo    string `json:"logo"`
}

type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Permissions string `json:"permissions"`
}

type UpdateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permissions string `json:"permissions"`
}

type CreateEmployeeRequest struct {
	NIK          string `json:"nik" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" binding:"required,email"`
	Phone        string `json:"phone"`
	DepartmentID uint   `json:"department_id" binding:"required"`
	LocationID   uint   `json:"location_id" binding:"required"`
	Position     string `json:"position"`
	Employment   string `json:"employment"`
	JoinDate     string `json:"join_date"`
}

type UpdateEmployeeRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	DepartmentID uint   `json:"department_id"`
	LocationID   uint   `json:"location_id"`
	Position     string `json:"position"`
	Employment   string `json:"employment"`
	Status       string `json:"status"`
	Address      string `json:"address"`
}

type CreateDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type UpdateDepartmentRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreateLocationRequest struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Timezone string `json:"timezone"`
}

type UpdateLocationRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Timezone string `json:"timezone"`
}
