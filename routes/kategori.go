package routes

import (
	"pelaporan_keuangan/features/kategori"

	"github.com/gin-gonic/gin"
)

func Kategori(r *gin.Engine, handler kategori.Handler) {
	kategori := r.Group("/kategori")

	kategori.GET("", handler.GetKategori)
	kategori.POST("", handler.CreateKategori)

	kategori.GET("/:id", handler.KategoriDetails)
	kategori.PUT("/:id", handler.UpdateKategori)
	kategori.DELETE("/:id", handler.DeleteKategori)
}
