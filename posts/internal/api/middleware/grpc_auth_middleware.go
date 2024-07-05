package middleware

import (
	"fmt"
	"net/http"
	"posts/internal/clients/auth/grpc"
	"posts/internal/schemas"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(client *grpc.GRPCClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")
		if len(token) == 2 {
			authToken := token[1]
			userID, err := client.ExtractUserIDFromToken(ctx, authToken)
			fmt.Printf("userID ==================== %v\nerrr ============ %v\n", userID, err)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
				ctx.Abort()
				return
			}
			ctx.Set("x-user-id", userID)
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		ctx.Abort()
	}
}
