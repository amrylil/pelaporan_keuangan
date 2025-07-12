package routes

import (
	"pelaporan_keuangan/features/users"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine, handler users.Handler) {
	auth := r.Group("/api/v1/auth")

	auth.POST("/login", handler.Login)
	auth.POST("/register", handler.CreateUser)
}
