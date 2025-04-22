package routes

import (
	"pelaporan_keuangan/features/master_data"

	"github.com/gin-gonic/gin"
)

func Master_data(r *gin.Engine, handler master_data.Handler) {
	master_data := r.Group("/master_data")

	master_data.GET("", handler.GetMaster_data)
	master_data.POST("", handler.CreateMaster_data)

	master_data.GET("/:id", handler.Master_dataDetails)
	master_data.PUT("/:id", handler.UpdateMaster_data)
	master_data.DELETE("/:id", handler.DeleteMaster_data)
}
