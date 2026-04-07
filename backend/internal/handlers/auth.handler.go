package handlers

import (
	"net/http"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *services.AuthService
}

func NewAuthHandler(userService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

// CreateNewUser godoc
// @Summary      Register user
// @Description  Creates a new user account
// @Tags         users
// @Accept json
// @Produce json
// @Param body body dto.RegisterDTO true "Registration payload"
// @Success 201 {object} dto.Response{data=dto.RegisterResponseDTO}
// @Failure      400 {object} dto.Response
// @Failure      409 {object} dto.Response
// @Failure      500 {object} dto.Response
// @Router       /auth/register [post]
func (u AuthHandler) Register(ctx *gin.Context) {
	var newUserData dto.RegisterDTO
	err := ctx.ShouldBind(&newUserData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: "Register failed! " + err.Error(),
			Data:    nil,
		})
		return
	}
	registeredUser, err := u.userService.Register(newUserData)
	if err != nil {
		if strings.Contains(err.Error(), "allready used") {
			ctx.JSON(http.StatusConflict, dto.ResponseDTO{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Internal server error",
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

func (u AuthHandler) Login(ctx *gin.Context){
	
}