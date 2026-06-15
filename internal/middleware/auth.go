package middleware

import (
	"strings"

	"hris-BE/pkg/jwt"
	"hris-BE/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtMgr *jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "Authorization header required")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "Invalid authorization format")
			c.Abort()
			return
		}

		claims, err := jwtMgr.Verify(parts[1])
		if err != nil {
			response.Unauthorized(c, "Invalid or expired token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("company_id", claims.CompanyID)
		c.Next()
	}
}

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			response.Unauthorized(c, "Role not found")
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			response.Unauthorized(c, "Invalid role")
			c.Abort()
			return
		}

		for _, r := range roles {
			if r == roleStr {
				c.Next()
				return
			}
		}

		response.Forbidden(c, "Insufficient permissions")
		c.Abort()
	}
}
