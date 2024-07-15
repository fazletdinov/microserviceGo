package route

import (
	"api-grpc-gateway/config"

	controller "api-grpc-gateway/internal/api/controller/auth"
	"api-grpc-gateway/internal/clients/auth"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(group *gin.RouterGroup, client *auth.GRPCClientAuth, env *config.Config) {
	profileController := &controller.ProfileController{
		GRPCClientAuth: client,
	}
	group.GET("/user/me", profileController.Fetch)
}
