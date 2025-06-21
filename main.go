package main

import (
	"fmt"
	"log"
	"net/http"
	"pelaporan_keuangan/config"
	"pelaporan_keuangan/features/kategori"
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

	kh "pelaporan_keuangan/features/kategori/handler"
	kr "pelaporan_keuangan/features/kategori/repository"
	ku "pelaporan_keuangan/features/kategori/usecase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	// Import generated docs - PENTING: Ini yang mungkin missing
	_ "pelaporan_keuangan/docs"
)

// @title Pelaporan Keuangan API
// @version 1.0
// @description API untuk sistem pelaporan keuangan dengan manajemen users, transaksi, dan master data
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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

	// Setup routes
	routes.Users(r, UsersHandler(db))
	routes.Transaksi(r, TransaksiHandler(db))
	routes.Master_data(r, MasterDataHandler(db))
	routes.Kategori(r, KategoriHandler(db))
	log.Println("Routes setup complete")

	// Root endpoint
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pelaporan Keuangan API is running! 😍")
	})

	// Health check endpoint
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Pelaporan Keuangan API is healthy",
		})
	})

	// Swagger endpoint - TAMBAHAN INI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Alternative swagger route (optional)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Swagger UI available at: http://localhost:" + cfg.SERVER_PORT + "/swagger/index.html")

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
func KategoriHandler(db *gorm.DB) kategori.Handler {
	repo := kr.New(db)
	usecase := ku.New(repo)
	return kh.New(usecase)
}
