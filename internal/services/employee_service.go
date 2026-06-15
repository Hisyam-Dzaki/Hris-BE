package services

import (
	"errors"
	"time"

	"hris-BE/internal/dto"
	"hris-BE/internal/models"
	"hris-BE/internal/repositories"
)

type EmployeeService struct {
	employeeRepo *repositories.EmployeeRepository
}

func NewEmployeeService(repo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{employeeRepo: repo}
}

func (s *EmployeeService) Create(companyID uint, req dto.CreateEmployeeRequest) (*models.Employee, error) {
	existing, _ := s.employeeRepo.FindByNIK(companyID, req.NIK)
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("NIK already exists")
	}

	employee := &models.Employee{
		CompanyID:    companyID,
		NIK:          req.NIK,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		Phone:        req.Phone,
		DepartmentID: req.DepartmentID,
		LocationID:   req.LocationID,
		Position:     req.Position,
		Employment:   req.Employment,
		Status:       "active",
	}

	if req.JoinDate != "" {
		t, err := time.Parse("2006-01-02", req.JoinDate)
		if err == nil {
			employee.JoinDate = &t
		}
	}

	if err := s.employeeRepo.Create(employee); err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *EmployeeService) GetByID(id uint) (*models.Employee, error) {
	return s.employeeRepo.FindByID(id)
}

func (s *EmployeeService) GetAll(filter repositories.EmployeeFilter) (*repositories.PaginatedEmployees, error) {
	return s.employeeRepo.FindAll(filter)
}

func (s *EmployeeService) Update(id uint, req dto.UpdateEmployeeRequest) (*models.Employee, error) {
	employee, err := s.employeeRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" {
		employee.FirstName = req.FirstName
	}
	if req.LastName != "" {
		employee.LastName = req.LastName
	}
	if req.Email != "" {
		employee.Email = req.Email
	}
	if req.Phone != "" {
		employee.Phone = req.Phone
	}
	if req.DepartmentID > 0 {
		employee.DepartmentID = req.DepartmentID
	}
	if req.LocationID > 0 {
		employee.LocationID = req.LocationID
	}
	if req.Position != "" {
		employee.Position = req.Position
	}
	if req.Employment != "" {
		employee.Employment = req.Employment
	}
	if req.Status != "" {
		employee.Status = req.Status
	}
	if req.Address != "" {
		employee.Address = req.Address
	}

	if err := s.employeeRepo.Update(employee); err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *EmployeeService) Delete(id uint) error {
	return s.employeeRepo.Delete(id)
}
