package handlers

import (
	"net/http"

	"hris-BE/internal/dto"
	"hris-BE/internal/services"
	"hris-BE/pkg/response"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	companyService *services.CompanyService
}

func NewCompanyHandler(svc *services.CompanyService) *CompanyHandler {
	return &CompanyHandler{companyService: svc}
}

func (h *CompanyHandler) Create(c *gin.Context) {
	var req dto.CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	company, err := h.companyService.Create(req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Company created", company)
}

func (h *CompanyHandler) GetByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	company, err := h.companyService.GetByID(id)
	if err != nil {
		response.NotFound(c, "Company not found")
		return
	}

	response.Success(c, http.StatusOK, "Success", company)
}

func (h *CompanyHandler) GetAll(c *gin.Context) {
	companies, err := h.companyService.GetAll()
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Success", companies)
}

func (h *CompanyHandler) Update(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req dto.UpdateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	company, err := h.companyService.Update(id, req)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Company updated", company)
}

func (h *CompanyHandler) Delete(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.companyService.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Company deleted", nil)
}
