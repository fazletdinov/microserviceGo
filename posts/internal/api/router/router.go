package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"posts/config"
)

func SetupRouter(env *config.Config, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/api/v1")
	NewCreatePostRouter(env, db, publicRouter)
	NewGetPostRouter(env, db, publicRouter)
	NewUpdatePostRouter(env, db, publicRouter)
	NewDeletePostRouter(env, db, publicRouter)
}
