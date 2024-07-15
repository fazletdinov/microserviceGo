package route

import (
	"api-grpc-gateway/config"

	controller "api-grpc-gateway/internal/api/controller/auth"
	"api-grpc-gateway/internal/clients/auth"

	"github.com/gin-gonic/gin"
)

func NewUpdateRouter(group *gin.RouterGroup, client *auth.GRPCClientAuth, env *config.Config) {
	updateController := &controller.UpdateController{
		GRPCClientAuth: client,
	}
	group.PUT("/user/update", updateController.Update)
}
