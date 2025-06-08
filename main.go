package main

import (
	"fmt"
	"log"
	"net/http"
	"pelaporan_keuangan/config"
	"pelaporan_keuangan/features/master_data"
	"pelaporan_keuangan/features/transaksi"
	"pelaporan_keuangan/features/users"
	"pelaporan_keuangan/features/users/handler"
	"pelaporan_keuangan/features/users/repository"
	"pelaporan_keuangan/features/users/usecase"
	"pelaporan_keuangan/routes"
	"pelaporan_keuangan/utils"

	mh "pelaporan_keuangan/features/master_data/handler"
	mr "pelaporan_keuangan/features/master_data/repository"
	mu "pelaporan_keuangan/features/master_data/usecase"

	th "pelaporan_keuangan/features/transaksi/handler"
	tr "pelaporan_keuangan/features/transaksi/repository"
	tu "pelaporan_keuangan/features/transaksi/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	log.Println("Starting application...")

	r := gin.Default()

	cfg := config.InitConfig()
	log.Printf("Config loaded: server port = %s", cfg.SERVER_PORT)

	db := utils.InitDB()
	log.Println("Database initialized")

	err := db.AutoMigrate(
		&users.Users{},
		&transaksi.Transaksi{},
		&master_data.JenisPembayaran{},
		&master_data.StatusTransaksi{},
		&master_data.TipeTransaksi{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration success")

	routes.Users(r, UsersHandler(db))
	routes.Transaksi(r, TransaksiHandler(db))
	routes.Master_data(r, MasterDataHandler(db))
	log.Println("Routes setup complete")

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello!😍")
	})

	log.Printf("Starting server on port :%s", cfg.SERVER_PORT)
	err = r.Run(fmt.Sprintf(":%s", cfg.SERVER_PORT))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func UsersHandler(db *gorm.DB) users.Handler {
	repo := repository.New(db)
	usecase := usecase.New(repo)
	return handler.New(usecase)
}

func TransaksiHandler(db *gorm.DB) transaksi.Handler {
	repo := tr.New(db)
	usecase := tu.New(repo)
	return th.New(usecase)
}

func MasterDataHandler(db *gorm.DB) master_data.Handler {
	repo := mr.New(db)
	usecase := mu.New(repo)
	return mh.New(usecase)
}
