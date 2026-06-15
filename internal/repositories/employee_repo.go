package repositories

import (
	"fmt"

	"hris-BE/internal/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *EmployeeRepository) FindByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.
		Preload("Company").
		Preload("Department").
		Preload("Location").
		Preload("User").
		First(&employee, id).Error
	return &employee, err
}

type EmployeeFilter struct {
	CompanyID    uint
	DepartmentID uint
	LocationID   uint
	Status       string
	Search       string
	Page         int
	Limit        int
}

type PaginatedEmployees struct {
	Employees []models.Employee
	Total     int64
	Page      int
	Limit     int
}

func (r *EmployeeRepository) FindAll(filter EmployeeFilter) (*PaginatedEmployees, error) {
	var employees []models.Employee
	var total int64

	query := r.db.Model(&models.Employee{}).
		Where("company_id = ?", filter.CompanyID).
		Preload("Department").
		Preload("Location")

	if filter.DepartmentID > 0 {
		query = query.Where("department_id = ?", filter.DepartmentID)
	}
	if filter.LocationID > 0 {
		query = query.Where("location_id = ?", filter.LocationID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR nik ILIKE ? OR email ILIKE ?",
			search, search, search, search)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("failed to count employees: %w", err)
	}

	offset := (filter.Page - 1) * filter.Limit
	if err := query.Offset(offset).Limit(filter.Limit).Order("id DESC").Find(&employees).Error; err != nil {
		return nil, fmt.Errorf("failed to find employees: %w", err)
	}

	return &PaginatedEmployees{
		Employees: employees,
		Total:     total,
		Page:      filter.Page,
		Limit:     filter.Limit,
	}, nil
}

func (r *EmployeeRepository) Update(employee *models.Employee) error {
	return r.db.Save(employee).Error
}

func (r *EmployeeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Employee{}, id).Error
}

func (r *EmployeeRepository) FindByNIK(companyID uint, nik string) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("company_id = ? AND nik = ?", companyID, nik).First(&employee).Error
	return &employee, err
}
