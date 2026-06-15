package handlers

import (
	"net/http"

	"hris-BE/internal/dto"
	"hris-BE/internal/services"
	"hris-BE/pkg/response"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	locService *services.LocationService
}

func NewLocationHandler(svc *services.LocationService) *LocationHandler {
	return &LocationHandler{locService: svc}
}

func (h *LocationHandler) Create(c *gin.Context) {
	var req dto.CreateLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	companyID := getCompanyID(c)
	loc, err := h.locService.Create(companyID, req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Location created", loc)
}

func (h *LocationHandler) GetByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	loc, err := h.locService.GetByID(id)
	if err != nil {
		response.NotFound(c, "Location not found")
		return
	}

	response.Success(c, http.StatusOK, "Success", loc)
}

func (h *LocationHandler) GetByCompanyID(c *gin.Context) {
	companyID := getCompanyID(c)
	locs, err := h.locService.GetByCompanyID(companyID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Success", locs)
}

func (h *LocationHandler) Update(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req dto.UpdateLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	loc, err := h.locService.Update(id, req)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Location updated", loc)
}

func (h *LocationHandler) Delete(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.locService.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Location deleted", nil)
}
