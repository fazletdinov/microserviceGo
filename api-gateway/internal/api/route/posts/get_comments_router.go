package route

import (
	"api-grpc-gateway/config"
	controller "api-grpc-gateway/internal/api/controller/posts"
	"api-grpc-gateway/internal/clients/posts"

	"github.com/gin-gonic/gin"
)

func NewGetCommentRouter(
	group *gin.RouterGroup,
	client *posts.GRPCClientPosts,
	env *config.Config,
) {
	postsController := &controller.GetCommentController{
		GRPCClientPosts: client,
		Env:             env,
	}
	group.GET("/post/:post_id/comments", postsController.Fetchs)
}
