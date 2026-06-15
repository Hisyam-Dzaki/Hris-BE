package services

import (
	"hris-BE/internal/dto"
	"hris-BE/internal/models"
	"hris-BE/internal/repositories"
)

type RoleService struct {
	roleRepo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{roleRepo: repo}
}

func (s *RoleService) Create(companyID uint, req dto.CreateRoleRequest) (*models.Role, error) {
	role := &models.Role{
		CompanyID:   companyID,
		Name:        req.Name,
		Description: req.Description,
		Permissions: req.Permissions,
	}

	if err := s.roleRepo.Create(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleService) GetByID(id uint) (*models.Role, error) {
	return s.roleRepo.FindByID(id)
}

func (s *RoleService) GetByCompanyID(companyID uint) ([]models.Role, error) {
	return s.roleRepo.FindByCompanyID(companyID)
}

func (s *RoleService) Update(id uint, req dto.UpdateRoleRequest) (*models.Role, error) {
	role, err := s.roleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		role.Name = req.Name
	}
	if req.Description != "" {
		role.Description = req.Description
	}
	if req.Permissions != "" {
		role.Permissions = req.Permissions
	}

	if err := s.roleRepo.Update(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleService) Delete(id uint) error {
	return s.roleRepo.Delete(id)
}
