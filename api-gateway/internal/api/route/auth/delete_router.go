package route

import (
	"api-grpc-gateway/config"

	controller "api-grpc-gateway/internal/api/controller/auth"
	"api-grpc-gateway/internal/clients/auth"

	"github.com/gin-gonic/gin"
)

func NewDeleteRouter(group *gin.RouterGroup, client *auth.GRPCClientAuth, env *config.Config) {
	deleteController := &controller.DeleteController{
		GRPCClientAuth: client,
		Env:            env,
	}
	group.DELETE("/user/delete", deleteController.Delete)
}
