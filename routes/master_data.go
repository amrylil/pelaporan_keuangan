package routes

import (
	"pelaporan_keuangan/features/master_data"

	"github.com/gin-gonic/gin"
)

func Master_data(r *gin.Engine, handler master_data.Handler) {
	masterData := r.Group("/master_data")

	// Jenis Pembayaran
	masterData.GET("/jenis-pembayaran", handler.GetJenisPembayaran)
	masterData.POST("/jenis-pembayaran", handler.CreateJenisPembayaran)
	masterData.GET("/jenis-pembayaran/:id", handler.JenisPembayaranDetails)
	masterData.PUT("/jenis-pembayaran/:id", handler.UpdateJenisPembayaran)
	masterData.DELETE("/jenis-pembayaran/:id", handler.DeleteJenisPembayaran)

	// Tipe Transaksi
	masterData.GET("/tipe-transaksi", handler.GetTipeTransaksi)
	masterData.POST("/tipe-transaksi", handler.CreateTipeTransaksi)
	masterData.GET("/tipe-transaksi/:id", handler.TipeTransaksiDetails)
	masterData.PUT("/tipe-transaksi/:id", handler.UpdateTipeTransaksi)
	masterData.DELETE("/tipe-transaksi/:id", handler.DeleteTipeTransaksi)

	// Status Transaksi
	masterData.GET("/status-transaksi", handler.GetStatusTransaksi)
	masterData.POST("/status-transaksi", handler.CreateStatusTransaksi)
	masterData.GET("/status-transaksi/:id", handler.StatusTransaksiDetails)
	masterData.PUT("/status-transaksi/:id", handler.UpdateStatusTransaksi)
	masterData.DELETE("/status-transaksi/:id", handler.DeleteStatusTransaksi)
}
