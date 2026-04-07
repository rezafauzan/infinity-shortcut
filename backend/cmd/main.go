package main

import (
	"fmt"
	"os"
	"snowfoxinfinity/infinity-shortcut/internal/di"
	"snowfoxinfinity/infinity-shortcut/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer recover()
	godotenv.Load()

	router := gin.Default()
	container, err := di.NewContainer()

	if err != nil {
		panic("Container Error : " + err.Error())
	}
	
	apiRouter := router.Group("/api")
	routers.NewAuthRouters(apiRouter, container)
	routers.NewLinkRouters(apiRouter, container)

	router.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}
