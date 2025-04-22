package routes

import (
	"pelaporan_keuangan/features/transaksi"

	"github.com/gin-gonic/gin"
)

func Transaksi(r *gin.Engine, handler transaksi.Handler) {
	transaksi := r.Group("/transaksi")

	transaksi.GET("", handler.GetTransaksi)
	transaksi.POST("", handler.CreateTransaksi)

	transaksi.GET("/:id", handler.TransaksiDetails)
	transaksi.PUT("/:id", handler.UpdateTransaksi)
	transaksi.DELETE("/:id", handler.DeleteTransaksi)
}
