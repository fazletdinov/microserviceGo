package route

import (
	"api-grpc-gateway/config"
	controller "api-grpc-gateway/internal/api/controller/likes"
	"api-grpc-gateway/internal/clients/likes"

	"github.com/gin-gonic/gin"
)

func NewDeleteReactionRouter(
	group *gin.RouterGroup,
	client *likes.GRPCClientLikes,
	env *config.Config,
) {
	likesController := &controller.DeleteReactionController{
		GRPCClientLikes: client,
		Env:             env,
	}

	group.DELETE("/post/:post_id/reaction", likesController.Delete)
}
