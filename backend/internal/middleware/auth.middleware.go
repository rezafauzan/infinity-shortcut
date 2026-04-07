package middleware

import (
	"net/http"
	"snowfoxinfinity/infinity-shortcut/internal/dto"
	"snowfoxinfinity/infinity-shortcut/internal/lib"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, dto.ResponseDTO{
				Success: false,
				Message: "Unauthorized access please login!",
				Data:    nil,
			})
			ctx.Abort()
			return
		}
		
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, dto.ResponseDTO{
				Success: false,
				Message: "Invalid token format",
				Data:    nil,
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := lib.VerifyJWT(tokenString)
		
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, dto.ResponseDTO{
				Success: false,
				Message: "Invalid or expired token",
				Data:    nil,
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", int(claims.UserId))
		ctx.Set("cart_id", int(claims.CartId))
		ctx.Set("role", "user")
		ctx.Next()
	}
}