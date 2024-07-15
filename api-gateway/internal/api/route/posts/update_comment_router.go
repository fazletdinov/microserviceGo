package route

import (
	"api-grpc-gateway/config"
	controller "api-grpc-gateway/internal/api/controller/posts"
	"api-grpc-gateway/internal/clients/posts"

	"github.com/gin-gonic/gin"
)

func NewUpdateCommentRouter(group *gin.RouterGroup, client *posts.GRPCClientPosts, env *config.Config) {
	postsController := &controller.UpdateCommentController{
		GRPCClientPosts: client,
		Env:             env,
	}
	group.PUT("/post/:post_id/comment/:comment_id", postsController.Update)
}
