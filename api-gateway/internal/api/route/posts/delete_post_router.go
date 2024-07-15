package route

import (
	"api-grpc-gateway/config"
	controller "api-grpc-gateway/internal/api/controller/posts"
	"api-grpc-gateway/internal/clients/posts"

	"github.com/gin-gonic/gin"
)

func NewDeletePostRouter(group *gin.RouterGroup, client *posts.GRPCClientPosts, env *config.Config) {
	postsController := &controller.DeletePostController{
		GRPCClientPosts: client,
		Env:             env,
	}

	group.DELETE("/post/:post_id", postsController.Delete)
}
