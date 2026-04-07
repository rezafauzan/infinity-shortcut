package handlers

import (
	"fmt"
	"net/http"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/services"
	"strconv"
	"strings"

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

// GetLinkById godoc
// @Summary      Get link by id
// @Description  Get link by id
// @Tags         links
// @Produce      json
// @Success      200 {object} dto.ResponseDTO{data=[]dto.GetAllLinksResponseDTO}
// @Failure      401 {object} dto.ResponseDTO
// @Failure      500 {object} dto.ResponseDTO
// @Router       /api/links/{id} [get]
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

// GetLinkBySlug godoc
// @Summary      Get link by slug
// @Description  Get link by slug
// @Tags         links
// @Produce      json
// @Success      200 {object} dto.ResponseDTO{data=[]dto.GetAllLinksResponseDTO}
// @Failure      401 {object} dto.ResponseDTO
// @Failure      500 {object} dto.ResponseDTO
// @Router       /api/links/{slug} [get]
func (l LinkHandler) GetLinkBySlug(ctx *gin.Context) {
	slugParam := ctx.Param("slug")
	links, err := l.linkService.GetLinkBySlug(slugParam)
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

// DeleteLinkById godoc
// @Summary      Delete link
// @Description  Delete link by id (only owner can delete)
// @Tags         links
// @Produce      json
// @Param        id   path      int  true  "Link ID"
// @Success      200  {object}  dto.ResponseDTO{data=dto.DeleteLinkResponseDTO}
// @Failure      400  {object}  dto.ResponseDTO
// @Failure      401  {object}  dto.ResponseDTO
// @Failure      403  {object}  dto.ResponseDTO
// @Failure      404  {object}  dto.ResponseDTO
// @Failure      500  {object}  dto.ResponseDTO
// @Router       /api/links/{id} [delete]
func (l LinkHandler) DeleteLinkById(ctx *gin.Context) {
	linkIdParam := ctx.Param("link_id")
	fmt.Println(linkIdParam)
	linkId, err := strconv.Atoi(linkIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Invalid id parameter!",
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

	deletedLink, err := l.linkService.DeleteLinkById(linkId, userId.(int))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.JSON(http.StatusNotFound, dto.ResponseDTO{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		if strings.Contains(err.Error(), "Forbidden") {
			ctx.JSON(http.StatusForbidden, dto.ResponseDTO{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Internal server error! " + err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Delete link success!",
		Data:    deletedLink,
	})
}
