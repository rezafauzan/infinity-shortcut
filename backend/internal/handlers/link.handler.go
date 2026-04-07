package handlers

import (
	"net/http"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	linkService *services.LinkService
}

func NewLinkHandler(linkService *services.LinkService) *LinkHandler {
	return &LinkHandler{
		linkService: linkService,
	}
}

// CreateNewLink godoc
// @Summary      Create new shortened link
// @Description  Create new shortened link
// @Tags         links
// @Accept json
// @Produce json
// @Param body body dto.CreateNewLinkDTO true "Registration payload"
// @Success 201 {object} dto.Response{data=dto.CreateNewLinkResponseDTO}
// @Failure      400 {object} dto.Response
// @Failure      500 {object} dto.Response
// @Router       /api/link [post]
func (l LinkHandler) CreateNewLink(ctx *gin.Context) {
	var newLinkData dto.CreateNewLinkDTO
	err := ctx.ShouldBind(&newLinkData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Create new link failed! " + err.Error(),
			Data:    nil,
		})
		return
	}

	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.ResponseDTO{
			Success: false,
			Message: "Unauthorized access please login!",
			Data:    nil,
		})
		return
	}

	registeredUser, err := l.linkService.CreateNewLink(newLinkData, userId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Internal server error" + err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.ResponseDTO{
		Success: true,
		Message: "Register success!",
		Data:    registeredUser,
	})
}

// GetAllLinksByUserId godoc
// @Summary      Get all links by user
// @Description  Get all links by user id from logged in user
// @Tags         links
// @Produce      json
// @Success      200 {object} dto.ResponseDTO{data=[]dto.GetAllLinksResponseDTO}
// @Failure      401 {object} dto.ResponseDTO
// @Failure      500 {object} dto.ResponseDTO
// @Router       /api/links [get]
func (l LinkHandler) GetLinkById(ctx *gin.Context) {
	linkIdParam := ctx.Param("link_id")

	linkId, err := strconv.Atoi(linkIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid link id",
			Data:    nil,
		})
		return
	}

	links, err := l.linkService.GetLinkById(linkId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Internal server error!",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Get all links success!",
		Data:    links,
	})
}

// GetAllLinksByUserId godoc
// @Summary      Get all links by user
// @Description  Get all links by user id from logged in user
// @Tags         links
// @Produce      json
// @Success      200 {object} dto.ResponseDTO{data=[]dto.GetAllLinksResponseDTO}
// @Failure      401 {object} dto.ResponseDTO
// @Failure      500 {object} dto.ResponseDTO
// @Router       /api/links [get]
func (l LinkHandler) GetAllLinksByUserId(ctx *gin.Context) {
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.ResponseDTO{
			Success: false,
			Message: "Unauthorized access please login!",
			Data:    nil,
		})
		return
	}

	links, err := l.linkService.GetAllLinksByUserId(userId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Internal server error!",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Get all links success!",
		Data:    links,
	})
}
