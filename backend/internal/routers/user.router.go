package routers

import (
	"snowfoxinfinity/infinity-shortcut/internal/di"

	"github.com/gin-gonic/gin"
)

func NewUserRouters(apiRouter *gin.RouterGroup, container *di.Container) {
	apiRouter.POST("register", container.UserHandler.CreateNewUser)
}