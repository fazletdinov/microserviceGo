package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"posts/config"
	commentRoute "posts/internal/api/router/comment"
	postRoute "posts/internal/api/router/post"
)

func SetupRouter(env *config.Config, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/api/v1")
	postRoute.NewCreatePostRouter(env, db, publicRouter)
	postRoute.NewGetPostRouter(env, db, publicRouter)
	postRoute.NewUpdatePostRouter(env, db, publicRouter)
	postRoute.NewDeletePostRouter(env, db, publicRouter)
	commentRoute.NewCreateCommentRouter(env, db, publicRouter)
	commentRoute.NewGetCommentRouter(env, db, publicRouter)
	commentRoute.NewUpdateCommentRouter(env, db, publicRouter)
	commentRoute.NewDeleteCommentRouter(env, db, publicRouter)
}
