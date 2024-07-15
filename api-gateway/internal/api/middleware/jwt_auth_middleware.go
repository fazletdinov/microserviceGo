package middleware

import (
	tokenService "api-grpc-gateway/internal/domain/services/auth"
	schemas "api-grpc-gateway/internal/schemas/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(pathSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")
		if len(token) == 2 {
			authToken := token[1]
			authorized, err := tokenService.IsAuthorized(authToken, pathSecret)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
				ctx.Abort()
				return
			}
			if authorized {
				userID, err := tokenService.ExtractUserIDFromToken(authToken, pathSecret)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
					ctx.Abort()
					return
				}
				ctx.Set("x-user-id", userID.String())
				ctx.Next()
				return
			}
			ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		ctx.Abort()
	}
}
