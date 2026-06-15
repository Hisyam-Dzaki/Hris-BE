package services

import (
	"hris-BE/internal/dto"
	"hris-BE/internal/models"
	"hris-BE/internal/repositories"
)

type DepartmentService struct {
	deptRepo *repositories.DepartmentRepository
}

func NewDepartmentService(repo *repositories.DepartmentRepository) *DepartmentService {
	return &DepartmentService{deptRepo: repo}
}

func (s *DepartmentService) Create(companyID uint, req dto.CreateDepartmentRequest) (*models.Department, error) {
	dept := &models.Department{
		CompanyID: companyID,
		Name:      req.Name,
		Code:      req.Code,
	}

	if err := s.deptRepo.Create(dept); err != nil {
		return nil, err
	}
	return dept, nil
}

func (s *DepartmentService) GetByID(id uint) (*models.Department, error) {
	return s.deptRepo.FindByID(id)
}

func (s *DepartmentService) GetByCompanyID(companyID uint) ([]models.Department, error) {
	return s.deptRepo.FindByCompanyID(companyID)
}

func (s *DepartmentService) Update(id uint, req dto.UpdateDepartmentRequest) (*models.Department, error) {
	dept, err := s.deptRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		dept.Name = req.Name
	}
	if req.Code != "" {
		dept.Code = req.Code
	}

	if err := s.deptRepo.Update(dept); err != nil {
		return nil, err
	}
	return dept, nil
}

func (s *DepartmentService) Delete(id uint) error {
	return s.deptRepo.Delete(id)
}
