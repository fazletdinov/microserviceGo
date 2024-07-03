package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"likes/config"
)

func SetupRouter(env *config.Config, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/api/v1")
	NewCreateReactionRouter(env, db, publicRouter)
	NewDeleteReactionRouter(env, db, publicRouter)
	NewGetReactionRouter(env, db, publicRouter)
}
