package route

import (
	"api-grpc-gateway/config"
	controller "api-grpc-gateway/internal/api/controller/likes"
	"api-grpc-gateway/internal/clients/likes"

	"github.com/gin-gonic/gin"
)

func NewCreateReactionRouter(group *gin.RouterGroup, client *likes.GRPCClientLikes, env *config.Config) {
	likesController := &controller.CreateReactionController{
		GRPCClientLikes: client,
		Env:             env,
	}
	group.POST("/post/:post_id/reaction", likesController.Create)
}
