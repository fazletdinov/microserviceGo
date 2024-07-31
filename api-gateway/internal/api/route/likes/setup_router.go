package route

import (
	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/api/middleware"
	"api-grpc-gateway/internal/clients/likes"

	"github.com/gin-gonic/gin"
)

func SetupLikesRouter(
	gin *gin.Engine,
	client *likes.GRPCClientLikes,
	env *config.Config,
) {
	publicRouter := gin.Group("/api/v1")
	NewGetReactionRouter(publicRouter, client, env)

	protectedRouter := gin.Group("/api/v1")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.JWTConfig.PathPublicKey))
	NewDeleteReactionRouter(protectedRouter, client, env)
	NewCreateReactionRouter(protectedRouter, client, env)
}
