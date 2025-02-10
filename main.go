package main

import (
	"blueprint_golang/config"
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

	r.Run(fmt.Sprintf(":%s", cfg.SERVER_PORT))
}
