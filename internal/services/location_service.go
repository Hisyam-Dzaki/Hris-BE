package services

import (
	"hris-BE/internal/dto"
	"hris-BE/internal/models"
	"hris-BE/internal/repositories"
)

type LocationService struct {
	locRepo *repositories.LocationRepository
}

func NewLocationService(repo *repositories.LocationRepository) *LocationService {
	return &LocationService{locRepo: repo}
}

func (s *LocationService) Create(companyID uint, req dto.CreateLocationRequest) (*models.Location, error) {
	loc := &models.Location{
		CompanyID: companyID,
		Name:      req.Name,
		Address:   req.Address,
		City:      req.City,
		Timezone:  req.Timezone,
	}

	if loc.Timezone == "" {
		loc.Timezone = "Asia/Jakarta"
	}

	if err := s.locRepo.Create(loc); err != nil {
		return nil, err
	}
	return loc, nil
}

func (s *LocationService) GetByID(id uint) (*models.Location, error) {
	return s.locRepo.FindByID(id)
}

func (s *LocationService) GetByCompanyID(companyID uint) ([]models.Location, error) {
	return s.locRepo.FindByCompanyID(companyID)
}

func (s *LocationService) Update(id uint, req dto.UpdateLocationRequest) (*models.Location, error) {
	loc, err := s.locRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		loc.Name = req.Name
	}
	if req.Address != "" {
		loc.Address = req.Address
	}
	if req.City != "" {
		loc.City = req.City
	}
	if req.Timezone != "" {
		loc.Timezone = req.Timezone
	}

	if err := s.locRepo.Update(loc); err != nil {
		return nil, err
	}
	return loc, nil
}

func (s *LocationService) Delete(id uint) error {
	return s.locRepo.Delete(id)
}
