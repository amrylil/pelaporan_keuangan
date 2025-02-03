package routes

import (
	"blueprint_golang/features/user"

	"github.com/gin-gonic/gin"
)

func Users(r *gin.Engine, handler user.Handler) {
	users := r.Group("/users")

	users.GET("", handler.GetUsers)
	users.POST("", handler.CreateUser)

	users.GET("/:id", handler.UserDetails)
	users.PUT("/:id", handler.UpdateUser)
	users.DELETE("/:id", handler.DeleteUser)
}
