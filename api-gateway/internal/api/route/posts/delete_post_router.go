package route

import (
	"api-grpc-gateway/config"
	controller "api-grpc-gateway/internal/api/controller/posts"
	"api-grpc-gateway/internal/clients/likes"
	"api-grpc-gateway/internal/clients/posts"

	"github.com/gin-gonic/gin"
)

func NewDeletePostRouter(
	group *gin.RouterGroup,
	clientPosts *posts.GRPCClientPosts,
	clientLikes *likes.GRPCClientLikes,
	env *config.Config,
) {
	postsController := &controller.DeletePostController{
		GRPCClientPosts: clientPosts,
		GRPCClientLikes: clientLikes,
		Env:             env,
	}

	group.DELETE("/post/:post_id", postsController.Delete)
}
