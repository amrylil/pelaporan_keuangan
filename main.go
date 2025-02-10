package main

import (
	"blueprint_golang/config"
	"blueprint_golang/features/product"
	"blueprint_golang/features/product/handler"
	"blueprint_golang/features/product/repository"
	"blueprint_golang/features/product/usecase"
	"blueprint_golang/routes"
	"blueprint_golang/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg := config.InitConfig()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello!😍")
	})

	routes.Products(r, ProductHandler())
	r.Run(fmt.Sprintf(":%s", cfg.SERVER_PORT))
}

func ProductHandler() product.Handler {
	db := utils.InitDB()
	db.AutoMigrate(product.Product{})

	repo := repository.New(db)
	usecase := usecase.New(repo)
	return handler.New(usecase)
}
