package routers

import (
	"snowfoxinfinity/infinity-shortcut/internal/di"

	"github.com/gin-gonic/gin"
)

func NewAuthRouters(apiRouter *gin.RouterGroup, container *di.Container) {
	apiRouter.POST("register", container.AuthHandler.Register)
	apiRouter.POST("login", container.AuthHandler.Login)
}
