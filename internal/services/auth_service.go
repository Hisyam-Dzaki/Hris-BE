package services

import (
	"errors"

	"hris-BE/internal/dto"
	"hris-BE/internal/models"
	"hris-BE/internal/repositories"
	bcryptpkg "hris-BE/pkg/bcrypt"
	jwtpkg "hris-BE/pkg/jwt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
	jwtMgr   *jwtpkg.JWTManager
}

func NewAuthService(userRepo *repositories.UserRepository, jwtMgr *jwtpkg.JWTManager) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtMgr:   jwtMgr,
	}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {
	existing, _ := s.userRepo.FindByEmail(req.Email)
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcryptpkg.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		CompanyID: req.CompanyID,
		Email:     req.Email,
		Password:  hashedPassword,
		Name:      req.Name,
		RoleName:  "employee",
		IsActive:  true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := s.jwtMgr.Generate(user.ID, user.Email, user.RoleName, user.CompanyID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !bcryptpkg.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	token, err := s.jwtMgr.Generate(user.ID, user.Email, user.RoleName, user.CompanyID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *AuthService) GetMe(userID uint) (*models.User, error) {
	return s.userRepo.FindByID(userID)
}
