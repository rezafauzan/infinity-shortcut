package routers

import (
	"snowfoxinfinity/infinity-shortcut/internal/di"
	"snowfoxinfinity/infinity-shortcut/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewLinkRouters(apiRouter *gin.RouterGroup, container *di.Container) {
	apiRouter.GET("links", middleware.AuthMiddleware(), container.LinkHandler.GetAllLinksByUserId)
	apiRouter.POST("links", middleware.AuthMiddleware(), container.LinkHandler.CreateNewLink)
}