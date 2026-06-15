package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseID(c *gin.Context, param string) (uint, error) {
	idStr := c.Param(param)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid %s", param)
	}
	return uint(id), nil
}

func getCompanyID(c *gin.Context) uint {
	cid, _ := c.Get("company_id")
	return cid.(uint)
}

func getUserID(c *gin.Context) uint {
	uid, _ := c.Get("user_id")
	return uid.(uint)
}
