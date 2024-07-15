package route

import (
	"api-grpc-gateway/config"
	controller "api-grpc-gateway/internal/api/controller/posts"
	"api-grpc-gateway/internal/clients/posts"

	"github.com/gin-gonic/gin"
)

func NewCreateCommentRouter(group *gin.RouterGroup, client *posts.GRPCClientPosts, env *config.Config) {
	postsController := &controller.CreateCommentController{
		GRPCClientPosts: client,
		Env:             env,
	}
	group.POST("/post/:post_id/comment", postsController.Create)
}
