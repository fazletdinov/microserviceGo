package route

import (
	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/api/middleware"
	"api-grpc-gateway/internal/clients/likes"
	"api-grpc-gateway/internal/clients/posts"

	"github.com/gin-gonic/gin"
)

func SetupPostsRouter(
	gin *gin.Engine,
	clientPosts *posts.GRPCClientPosts,
	clientLikes *likes.GRPCClientLikes,
	env *config.Config,
) {
	publicRouter := gin.Group("/api/v1")
	NewGetCommentRouter(publicRouter, clientPosts, env)
	NewGetPostRouter(publicRouter, clientPosts, env)

	protectedRouter := gin.Group("/api/v1")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.JWTConfig.PathPublicKey))
	NewCreatePostRouter(protectedRouter, clientPosts, env)
	NewCreateCommentRouter(protectedRouter, clientPosts, env)
	NewDeletePostRouter(protectedRouter, clientPosts, clientLikes, env)
	NewDeleteCommentRouter(protectedRouter, clientPosts, env)
	NewUpdateCommentRouter(protectedRouter, clientPosts, env)
	NewUpdatePostRouter(protectedRouter, clientPosts, env)
}
