package routes

import (
	"pelaporan_keuangan/features/auth"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine, handler auth.Handler) {
	auth := r.Group("/api/v1/auth")

	auth.GET("", handler.GetAuth)
	auth.POST("", handler.CreateAuth)

	auth.GET("/:id", handler.AuthDetails)
	auth.PUT("/:id", handler.UpdateAuth)
	auth.DELETE("/:id", handler.DeleteAuth)
}
