package route

import (
	"api-grpc-gateway/config"

	controller "api-grpc-gateway/internal/api/controller/auth"
	"api-grpc-gateway/internal/clients/auth"
	"api-grpc-gateway/internal/clients/likes"
	"api-grpc-gateway/internal/clients/posts"

	"github.com/gin-gonic/gin"
)

func NewDeleteRouter(
	group *gin.RouterGroup,
	clientAuth *auth.GRPCClientAuth,
	clientPost *posts.GRPCClientPosts,
	clientLikes *likes.GRPCClientLikes,
	env *config.Config,
) {
	deleteController := &controller.DeleteController{
		GRPCClientAuth:  clientAuth,
		GRPCClientPosts: clientPost,
		GRPCClientLikes: clientLikes,
		Env:             env,
	}
	group.DELETE("/user/delete", deleteController.Delete)
}
