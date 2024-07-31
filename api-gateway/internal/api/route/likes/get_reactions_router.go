package route

import (
	"api-grpc-gateway/config"
	controller "api-grpc-gateway/internal/api/controller/likes"
	"api-grpc-gateway/internal/clients/likes"

	"github.com/gin-gonic/gin"
)

func NewGetReactionRouter(
	group *gin.RouterGroup,
	client *likes.GRPCClientLikes,
	env *config.Config,
) {
	likesController := &controller.GetReactionController{
		GRPCClientLikes: client,
		Env:             env,
	}
	group.GET("/post/:post_id/reactions", likesController.Fetchs)
}
