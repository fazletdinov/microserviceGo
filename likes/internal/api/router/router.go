package router

import (
	"likes/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(env *config.Config, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/api/v1")
	NewCreateReactionRouter(env, db, publicRouter)
	NewDeleteReactionRouter(env, db, publicRouter)
	NewGetReactionRouter(env, db, publicRouter)
}
