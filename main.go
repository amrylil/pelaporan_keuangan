package main

import (
	"fmt"
	"net/http"
	"pelaporan_keuangan/config"
	"pelaporan_keuangan/features/master_data"
	"pelaporan_keuangan/features/users"
	"pelaporan_keuangan/features/users/handler"
	"pelaporan_keuangan/features/users/repository"
	"pelaporan_keuangan/features/users/usecase"
	"pelaporan_keuangan/routes"
	"pelaporan_keuangan/utils"

	mh "pelaporan_keuangan/features/users/handler"
	mr "pelaporan_keuangan/features/users/repository"
	mu "pelaporan_keuangan/features/users/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg := config.InitConfig()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello!😍")
	})

	routes.Users(r, UsersHandler())
	r.Run(fmt.Sprintf(":%s", cfg.SERVER_PORT))
}

func UsersHandler() users.Handler {
	db := utils.InitDB()
	db.AutoMigrate(users.Users{})

	repo := repository.New(db)
	usecase := usecase.New(repo)
	return handler.New(usecase)
}

func MasterDataHandler() users.Handler {
	db := utils.InitDB()
	db.AutoMigrate(master_data.JenisPembayaran{})
	db.AutoMigrate(master_data.StatusTransaksi{})
	db.AutoMigrate(master_data.TipeTransaksi{})

	repo := mr.New(db)
	usecase := mu.New(repo)
	return mh.New(usecase)
}
