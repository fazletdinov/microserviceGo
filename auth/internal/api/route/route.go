package route

import (
	"auth/config"
	"auth/internal/api/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(env *config.Config, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/api/v1")
	NewSignupRouter(env, db, publicRouter)
	NewLoginRouter(env, db, publicRouter)
	NewRefreshTokenRouter(env, db, publicRouter)

	protectedRouter := gin.Group("/api/v1")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.JWTConfig.PathPublicKey))
	NewProfileRouter(env, db, protectedRouter)
	NewUpdateRouter(env, db, protectedRouter)
	NewDeleteRouter(env, db, protectedRouter)
}
