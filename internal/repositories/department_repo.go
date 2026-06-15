package repositories

import (
	"hris-BE/internal/models"

	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (r *DepartmentRepository) Create(dept *models.Department) error {
	return r.db.Create(dept).Error
}

func (r *DepartmentRepository) FindByID(id uint) (*models.Department, error) {
	var dept models.Department
	err := r.db.First(&dept, id).Error
	return &dept, err
}

func (r *DepartmentRepository) FindByCompanyID(companyID uint) ([]models.Department, error) {
	var depts []models.Department
	err := r.db.Where("company_id = ?", companyID).Order("name ASC").Find(&depts).Error
	return depts, err
}

func (r *DepartmentRepository) Update(dept *models.Department) error {
	return r.db.Save(dept).Error
}

func (r *DepartmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Department{}, id).Error
}
