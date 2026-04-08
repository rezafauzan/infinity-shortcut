package routers

import (
	"snowfoxinfinity/infinity-shortcut/internal/di"
	"snowfoxinfinity/infinity-shortcut/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewAuthRouters(apiRouter *gin.RouterGroup, container *di.Container) {
	apiRouter.POST("register", container.AuthHandler.Register)
	apiRouter.POST("login", container.AuthHandler.Login)
	apiRouter.GET("validate-token", middleware.AuthMiddleware(), container.AuthHandler.GetUserById)
}
