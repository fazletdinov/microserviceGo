package route

import (
	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/api/middleware"
	"api-grpc-gateway/internal/clients/auth"
	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(gin *gin.Engine, client *auth.GRPCClientAuth, env *config.Config) {
	publicRouter := gin.Group("/api/v1")
	NewSignupRouter(publicRouter, client, env)
	NewLoginRouter(publicRouter, client, env)
	NewRefreshRouter(publicRouter, client, env)

	protectedRouter := gin.Group("/api/v1")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.JWTConfig.PathPublicKey))
	NewDeleteRouter(protectedRouter, client, env)
	NewProfileRouter(protectedRouter, client, env)
	NewUpdateRouter(protectedRouter, client, env)
}
