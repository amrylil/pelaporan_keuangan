package routes

import (
	"blueprint_golang/features/product"

	"github.com/gin-gonic/gin"
)

func Products(r *gin.Engine, handler product.Handler) {
	products := r.Group("/products")

	products.GET("", handler.GetProducts)
	products.POST("", handler.CreateProduct)

	products.GET("/:id", handler.ProductDetails)
	products.PUT("/:id", handler.UpdateProduct)
	products.DELETE("/:id", handler.DeleteProduct)
}
