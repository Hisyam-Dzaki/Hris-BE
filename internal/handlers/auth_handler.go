package handlers

import (
	"net/http"

	"hris-BE/internal/dto"
	"hris-BE/internal/services"
	"hris-BE/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.authService.Register(req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Registration successful", result)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.authService.Login(req)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Login successful", result)
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "User not found")
		return
	}

	user, err := h.authService.GetMe(userID.(uint))
	if err != nil {
		response.NotFound(c, "User not found")
		return
	}

	response.Success(c, http.StatusOK, "Success", user)
}
