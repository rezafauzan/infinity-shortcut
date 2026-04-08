package handlers

import (
	"net/http"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
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
// @Router       /api/register [post]
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
	registeredUser, err := u.authService.Register(newUserData)
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

// Login godoc
// @Summary      Authenticate user
// @Description  Validates credentials and returns a JWT access token.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      dto.LoginRequestDTO  true  "Login credentials"
// @Success      200   {object}  dto.Response{data=dto.LoginResponseDTO}
// @Failure      400   {object}  dto.Response
// @Failure      401   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /api/login [post]
func (u AuthHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequestDTO
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	result, err := u.authService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseDTO{
			Success: false,
			Message: "Login failed " + err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.ResponseDTO{
		Success: true,
		Message: "Login Success!",
		Data:    result,
	})
}

// GetUserById godoc
// @Summary      Get user by id
// @Description  Get user detail from logged in user
// @Tags         user
// @Produce      json
// @Success      200 {object} dto.ResponseDTO{data=dto.GetUserResponseDTO}
// @Failure      401 {object} dto.ResponseDTO
// @Failure      500 {object} dto.ResponseDTO
// @Router       /api/user [get]
func (a AuthHandler) GetUserById(ctx *gin.Context) {
	userId, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, dto.ResponseDTO{
			Success: false,
			Message: "Unauthorized access please login!",
			Data:    nil,
		})
		return
	}

	user, err := a.authService.GetUserById(userId.(int))
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
		Message: "Get user success!",
		Data:    user,
	})
}
