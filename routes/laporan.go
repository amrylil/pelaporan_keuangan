package routes

import (
	"pelaporan_keuangan/features/laporan"

	"github.com/gin-gonic/gin"
)

func Laporan(r *gin.Engine, handler laporan.Handler) {
	laporan := r.Group("/laporan")

	laporan.GET("", handler.GetLaporan)
	laporan.POST("", handler.CreateLaporan)

	laporan.GET("/:id", handler.LaporanDetails)
	laporan.PUT("/:id", handler.UpdateLaporan)
	laporan.DELETE("/:id", handler.DeleteLaporan)
}
