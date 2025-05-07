package middleware

import (
	"errors"
	"net/http"
	"pelaporan_keuangan/pkg/jwt"
	"pelaporan_keuangan/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth is a middleware that validates JWT tokens
func JWTAuth(jwtManager *jwt.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractTokenFromHeader(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("unauthorized", err.Error()))
			return
		}

		claims, err := jwtManager.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("unauthorized", "Invalid token: "+err.Error()))
			return
		}

		// Set claims data to context for use in handlers
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("roles", claims.Roles)

		c.Next()
	}
}

// extractTokenFromHeader extracts the JWT token from Authorization header
func extractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("authorization header format must be Bearer {token}")
	}

	return parts[1], nil
}
