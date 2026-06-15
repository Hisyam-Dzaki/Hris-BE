package repositories

import (
	"hris-BE/internal/models"

	"gorm.io/gorm"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

func (r *LocationRepository) Create(loc *models.Location) error {
	return r.db.Create(loc).Error
}

func (r *LocationRepository) FindByID(id uint) (*models.Location, error) {
	var loc models.Location
	err := r.db.First(&loc, id).Error
	return &loc, err
}

func (r *LocationRepository) FindByCompanyID(companyID uint) ([]models.Location, error) {
	var locs []models.Location
	err := r.db.Where("company_id = ?", companyID).Order("name ASC").Find(&locs).Error
	return locs, err
}

func (r *LocationRepository) Update(loc *models.Location) error {
	return r.db.Save(loc).Error
}

func (r *LocationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Location{}, id).Error
}
