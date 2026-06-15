package handlers

import (
	"net/http"
	"strconv"

	"hris-BE/internal/dto"
	"hris-BE/internal/repositories"
	"hris-BE/internal/services"
	"hris-BE/pkg/response"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeService *services.EmployeeService
}

func NewEmployeeHandler(svc *services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{employeeService: svc}
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	companyID := getCompanyID(c)
	employee, err := h.employeeService.Create(companyID, req)
	if err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Employee created", employee)
}

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	employee, err := h.employeeService.GetByID(id)
	if err != nil {
		response.NotFound(c, "Employee not found")
		return
	}

	response.Success(c, http.StatusOK, "Success", employee)
}

func (h *EmployeeHandler) GetAll(c *gin.Context) {
	companyID := getCompanyID(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	deptID, _ := strconv.ParseUint(c.Query("department_id"), 10, 32)
	locID, _ := strconv.ParseUint(c.Query("location_id"), 10, 32)

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	filter := repositories.EmployeeFilter{
		CompanyID:    companyID,
		DepartmentID: uint(deptID),
		LocationID:   uint(locID),
		Status:       c.Query("status"),
		Search:       c.Query("search"),
		Page:         page,
		Limit:        limit,
	}

	result, err := h.employeeService.GetAll(filter)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	totalPages := int(result.Total) / limit
	if int(result.Total)%limit > 0 {
		totalPages++
	}

	response.SuccessWithMeta(c, http.StatusOK, result.Employees, &response.Meta{
		Page:       page,
		Limit:      limit,
		Total:      result.Total,
		TotalPages: totalPages,
	})
}

func (h *EmployeeHandler) Update(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	employee, err := h.employeeService.Update(id, req)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Employee updated", employee)
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.employeeService.Delete(id); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Employee deleted", nil)
}
