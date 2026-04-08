package main

import (
	"fmt"
	"net/http"
	"os"
	"snowfoxinfinity/infinity-shortcut/internal/di"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/middleware"
	"snowfoxinfinity/infinity-shortcut/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer recover()
	godotenv.Load()

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dto.ResponseDTO{
			Success: true,
			Message: "Backend running well !",
			Data: nil,
		})
	})
	container, err := di.NewContainer()

	if err != nil {
		panic("Container Error : " + err.Error())
	}

	router.Use(middleware.CORSMiddleware())

	apiRouter := router.Group("/api")
	routers.NewAuthRouters(apiRouter, container)
	routers.NewLinkRouters(apiRouter, container)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
