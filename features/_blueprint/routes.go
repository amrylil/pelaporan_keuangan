package routes

import (
	"pelaporan_keuangan/features/_blueprint"

	"github.com/gin-gonic/gin"
)

func Placeholder(r *gin.Engine, handler _blueprint.Handler) {
	placeholder := r.Group("/placeholder")

	placeholder.GET("", handler.GetPlaceholder)
	placeholder.POST("", handler.CreatePlaceholder)

	placeholder.GET("/:id", handler.PlaceholderDetails)
	placeholder.PUT("/:id", handler.UpdatePlaceholder)
	placeholder.DELETE("/:id", handler.DeletePlaceholder)
}
