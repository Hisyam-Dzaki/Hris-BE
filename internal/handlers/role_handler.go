package handlers

import (
	"net/http"

	"hris-BE/internal/dto"
	"hris-BE/internal/services"
	"hris-BE/pkg/response"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(svc *services.RoleService) *RoleHandler {
	return &RoleHandler{roleService: svc}
}

func (h *RoleHandler) Create(c *gin.Context) {
	var req dto.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	companyID := getCompanyID(c)
	role, err := h.roleService.Create(companyID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Role created", role)
}

func (h *RoleHandler) GetByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	role, err := h.roleService.GetByID(id)
	if err != nil {
		response.NotFound(c, "Role not found")
		return
	}

	response.Success(c, http.StatusOK, "Success", role)
}

func (h *RoleHandler) GetByCompanyID(c *gin.Context) {
	companyID := getCompanyID(c)
	roles, err := h.roleService.GetByCompanyID(companyID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Success", roles)
}

func (h *RoleHandler) Update(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req dto.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	role, err := h.roleService.Update(id, req)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Role updated", role)
}

func (h *RoleHandler) Delete(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.roleService.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Role deleted", nil)
}
