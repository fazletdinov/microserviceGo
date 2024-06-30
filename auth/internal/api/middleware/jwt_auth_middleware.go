package middleware

import (
	"auth/internal/domain/service"
	"auth/internal/schemas"
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
			authorized, err := service.IsAuthorized(authToken, pathSecret)
			if authorized {
				userID, err := service.ExtractUserIDFromToken(authToken, pathSecret)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: err.Error()})
					ctx.Abort()
					return
				}
				ctx.Set("x-user-id", userID.String())
				ctx.Next()
				return
			}
			ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: err.Error()})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Not authorized"})
		ctx.Abort()
	}
}
