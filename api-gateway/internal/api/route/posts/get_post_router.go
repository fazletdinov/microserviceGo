package route

import (
	"api-grpc-gateway/config"
	controller "api-grpc-gateway/internal/api/controller/posts"
	"api-grpc-gateway/internal/clients/posts"

	"github.com/gin-gonic/gin"
)

func NewGetPostRouter(group *gin.RouterGroup, client *posts.GRPCClientPosts, env *config.Config) {
	postsController := &controller.GetPostController{
		GRPCClientPosts: client,
		Env:             env,
	}
	group.GET("/post/:post_id", postsController.Fetch)
	group.GET("/posts", postsController.Fetchs)
}
