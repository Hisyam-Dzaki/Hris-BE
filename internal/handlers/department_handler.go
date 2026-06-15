package handlers

import (
	"net/http"

	"hris-BE/internal/dto"
	"hris-BE/internal/services"
	"hris-BE/pkg/response"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	deptService *services.DepartmentService
}

func NewDepartmentHandler(svc *services.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{deptService: svc}
}

func (h *DepartmentHandler) Create(c *gin.Context) {
	var req dto.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	companyID := getCompanyID(c)
	dept, err := h.deptService.Create(companyID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Department created", dept)
}

func (h *DepartmentHandler) GetByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	dept, err := h.deptService.GetByID(id)
	if err != nil {
		response.NotFound(c, "Department not found")
		return
	}

	response.Success(c, http.StatusOK, "Success", dept)
}

func (h *DepartmentHandler) GetByCompanyID(c *gin.Context) {
	companyID := getCompanyID(c)
	depts, err := h.deptService.GetByCompanyID(companyID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Success", depts)
}

func (h *DepartmentHandler) Update(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req dto.UpdateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	dept, err := h.deptService.Update(id, req)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Department updated", dept)
}

func (h *DepartmentHandler) Delete(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.deptService.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Department deleted", nil)
}
