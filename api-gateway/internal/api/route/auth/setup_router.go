package route

import (
	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/api/middleware"
	"api-grpc-gateway/internal/clients/auth"
	"api-grpc-gateway/internal/clients/likes"
	"api-grpc-gateway/internal/clients/posts"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(
	gin *gin.Engine,
	clientAuth *auth.GRPCClientAuth,
	clientPost *posts.GRPCClientPosts,
	clientLikes *likes.GRPCClientLikes,
	env *config.Config,
) {
	publicRouter := gin.Group("/api/v1")
	NewSignupRouter(publicRouter, clientAuth, env)
	NewLoginRouter(publicRouter, clientAuth, env)
	NewRefreshRouter(publicRouter, clientAuth, env)

	protectedRouter := gin.Group("/api/v1")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.JWTConfig.PathPublicKey))
	NewDeleteRouter(protectedRouter, clientAuth, clientPost, clientLikes, env)
	NewProfileRouter(protectedRouter, clientAuth, env)
	NewUpdateRouter(protectedRouter, clientAuth, env)
}
