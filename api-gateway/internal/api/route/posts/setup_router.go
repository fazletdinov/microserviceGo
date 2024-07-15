package route

import (
	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/api/middleware"
	"api-grpc-gateway/internal/clients/posts"
	"github.com/gin-gonic/gin"
)

func SetupPostsRouter(gin *gin.Engine, client *posts.GRPCClientPosts, env *config.Config) {
	publicRouter := gin.Group("/api/v1")
	NewGetCommentRouter(publicRouter, client, env)
	NewGetPostRouter(publicRouter, client, env)

	protectedRouter := gin.Group("/api/v1")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.JWTConfig.PathPublicKey))
	NewCreatePostRouter(protectedRouter, client, env)
	NewCreateCommentRouter(protectedRouter, client, env)
	NewDeletePostRouter(protectedRouter, client, env)
	NewDeleteCommentRouter(protectedRouter, client, env)
	NewUpdateCommentRouter(protectedRouter, client, env)
	NewUpdatePostRouter(protectedRouter, client, env)
}
