package route

import (
	"api-grpc-gateway/config"
	controller "api-grpc-gateway/internal/api/controller/auth"
	"api-grpc-gateway/internal/clients/auth"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(group *gin.RouterGroup, client *auth.GRPCClientAuth, env *config.Config) {
	signupController := controller.SignupController{
		GRPCClientAuth: client,
		Env:            env,
	}
	group.POST("/user/signup", signupController.Signup)
}
