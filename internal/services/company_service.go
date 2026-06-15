package services

import (
	"hris-BE/internal/dto"
	"hris-BE/internal/models"
	"hris-BE/internal/repositories"
)

type CompanyService struct {
	companyRepo *repositories.CompanyRepository
}

func NewCompanyService(repo *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{companyRepo: repo}
}

func (s *CompanyService) Create(req dto.CreateCompanyRequest) (*models.Company, error) {
	company := &models.Company{
		Name:    req.Name,
		Slug:    req.Slug,
		Address: req.Address,
		Phone:   req.Phone,
		Email:   req.Email,
		Website: req.Website,
	}

	if err := s.companyRepo.Create(company); err != nil {
		return nil, err
	}
	return company, nil
}

func (s *CompanyService) GetByID(id uint) (*models.Company, error) {
	return s.companyRepo.FindByID(id)
}

func (s *CompanyService) GetAll() ([]models.Company, error) {
	return s.companyRepo.FindAll()
}

func (s *CompanyService) Update(id uint, req dto.UpdateCompanyRequest) (*models.Company, error) {
	company, err := s.companyRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		company.Name = req.Name
	}
	if req.Address != "" {
		company.Address = req.Address
	}
	if req.Phone != "" {
		company.Phone = req.Phone
	}
	if req.Email != "" {
		company.Email = req.Email
	}
	if req.Website != "" {
		company.Website = req.Website
	}
	if req.Logo != "" {
		company.Logo = req.Logo
	}

	if err := s.companyRepo.Update(company); err != nil {
		return nil, err
	}
	return company, nil
}

func (s *CompanyService) Delete(id uint) error {
	return s.companyRepo.Delete(id)
}
