package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Page       int   `json:"page,omitempty"`
	Limit      int   `json:"limit,omitempty"`
	Total      int64 `json:"total,omitempty"`
	TotalPages int   `json:"total_pages,omitempty"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

func Success(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func SuccessWithMeta(c *gin.Context, status int, data interface{}, meta *Meta) {
	c.JSON(status, Response{
		Status: "success",
		Data:   data,
		Meta:   meta,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Status:  "error",
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message)
}

func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, message)
}

func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

func InternalError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}
