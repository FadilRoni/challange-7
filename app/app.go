package app

import (
	"challange-7/config"
	"challange-7/repository"
	"challange-7/route"
	"challange-7/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApp() {
	repo := repository.NewRepo(config.PSQL.DB)
	app := service.NewService(repo)
	// server := handler.NewHttpServer(app)

	route.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")

	router.Run(fmt.Sprintf(":%s", port))
}
