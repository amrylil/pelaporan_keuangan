package middleware

import (
	"net/http"
	"pelaporan_keuangan/pkg/utils"

	"github.com/gin-gonic/gin"
)

// RoleAuth is a middleware that checks if user has required roles
func RoleAuth(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get roles from context (set by JWTAuth middleware)
		rolesInterface, exists := c.Get("roles")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("unauthorized", "user roles not found"))
			return
		}

		userRoles, ok := rolesInterface.([]string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse("error", "invalid roles data type"))
			return
		}

		// Check if user has any of the required roles
		hasRequiredRole := false
		for _, requiredRole := range requiredRoles {
			for _, userRole := range userRoles {
				if requiredRole == userRole {
					hasRequiredRole = true
					break
				}
			}
			if hasRequiredRole {
				break
			}
		}

		if !hasRequiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.ErrorResponse("forbidden", "insufficient permissions"))
			return
		}

		c.Next()
	}
}

// RequireAllRoles is a middleware that checks if user has all required roles
func RequireAllRoles(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get roles from context (set by JWTAuth middleware)
		rolesInterface, exists := c.Get("roles")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("unauthorized", "user roles not found"))
			return
		}

		userRoles, ok := rolesInterface.([]string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorResponse("error", "invalid roles data type"))
			return
		}

		// Convert user roles to map for O(1) lookup
		userRolesMap := make(map[string]bool)
		for _, role := range userRoles {
			userRolesMap[role] = true
		}

		// Check if user has all required roles
		for _, requiredRole := range requiredRoles {
			if !userRolesMap[requiredRole] {
				c.AbortWithStatusJSON(http.StatusForbidden, utils.ErrorResponse("forbidden", "insufficient permissions"))
				return
			}
		}

		c.Next()
	}
}
