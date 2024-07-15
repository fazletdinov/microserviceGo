package route

import (
	"api-grpc-gateway/config"

	controller "api-grpc-gateway/internal/api/controller/auth"
	"api-grpc-gateway/internal/clients/auth"

	"github.com/gin-gonic/gin"
)

func NewRefreshRouter(group *gin.RouterGroup, client *auth.GRPCClientAuth, env *config.Config) {
	refreshController := &controller.RefreshTokenController{
		GRPCClientAuth: client,
		Env:            env,
	}
	group.GET("/refresh", refreshController.RefreshToken)
}
