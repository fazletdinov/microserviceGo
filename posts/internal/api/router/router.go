package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"posts/config"
	"posts/internal/api/middleware"
	commentRoute "posts/internal/api/router/comment"
	postRoute "posts/internal/api/router/post"
	"posts/internal/clients/auth/grpc"
)

func SetupRouter(env *config.Config, db *gorm.DB, gin *gin.Engine, client *grpc.GRPCClient) {
	publicRouter := gin.Group("/api/v1")
	postRoute.NewGetPostRouter(env, db, publicRouter)
	commentRoute.NewGetCommentRouter(env, db, publicRouter)

	protectedRouter := gin.Group("/api/v1")
	protectedRouter.Use(middleware.JwtAuthMiddleware(client))
	postRoute.NewCreatePostRouter(env, db, protectedRouter)
	commentRoute.NewCreateCommentRouter(env, db, protectedRouter)
	postRoute.NewUpdatePostRouter(env, db, protectedRouter)
	postRoute.NewDeletePostRouter(env, db, protectedRouter)
	commentRoute.NewUpdateCommentRouter(env, db, protectedRouter)
	commentRoute.NewDeleteCommentRouter(env, db, protectedRouter)
}
