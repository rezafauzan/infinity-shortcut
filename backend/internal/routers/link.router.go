package routers

import (
	"snowfoxinfinity/infinity-shortcut/internal/di"

	"github.com/gin-gonic/gin"
)

func NewLinkRouters(apiRouter *gin.RouterGroup, container *di.Container) {
	apiRouter.POST("links", container.LinkHandler.CreateNewLink)
}