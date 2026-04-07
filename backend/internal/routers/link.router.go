package routers

import (
	"snowfoxinfinity/infinity-shortcut/internal/di"
	"snowfoxinfinity/infinity-shortcut/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewLinkRouters(apiRouter *gin.RouterGroup, container *di.Container) {
	apiRouter.GET("links", middleware.AuthMiddleware(), container.LinkHandler.GetAllLinksByUserId)
	apiRouter.GET("links/:slug", container.LinkHandler.GetLinkBySlug)
	apiRouter.POST("links", middleware.AuthMiddleware(), container.LinkHandler.CreateNewLink)
	apiRouter.DELETE("links/:link_id", middleware.AuthMiddleware(), container.LinkHandler.DeleteLinkById)
}