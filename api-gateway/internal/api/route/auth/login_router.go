package route

import (
	"api-grpc-gateway/config"

	controller "api-grpc-gateway/internal/api/controller/auth"
	"api-grpc-gateway/internal/clients/auth"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(group *gin.RouterGroup, client *auth.GRPCClientAuth, env *config.Config) {
	loginController := &controller.LoginController{
		GRPCClientAuth: client,
		Env:            env,
	}
	group.POST("/login", loginController.Login)
}
