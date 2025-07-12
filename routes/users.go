package routes

import (
	"pelaporan_keuangan/features/users"

	"github.com/gin-gonic/gin"
)

func Users(r *gin.Engine, handler users.Handler) {
	users := r.Group("/api/v1/users")

	users.GET("", handler.GetUsers)
	users.POST("", handler.CreateUser)

	users.GET("/:id", handler.UserDetails)
	users.PUT("/:id", handler.UpdateUser)
	users.DELETE("/:id", handler.DeleteUser)

}
