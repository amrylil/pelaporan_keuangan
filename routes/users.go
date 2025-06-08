package routes

import (
	"pelaporan_keuangan/features/users"

	"github.com/gin-gonic/gin"
)

func Users(r *gin.Engine, handler users.Handler) {
	users := r.Group("/api/v1/users")

	users.GET("", handler.GetUsers)
	users.POST("", handler.CreateUsers)

	users.GET("/:id", handler.UsersDetails)
	users.PUT("/:id", handler.UpdateUsers)
	users.DELETE("/:id", handler.DeleteUsers)
}
