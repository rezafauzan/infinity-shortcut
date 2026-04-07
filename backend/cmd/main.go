package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer recover()
	godotenv.Load()

	router := gin.Default()
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
