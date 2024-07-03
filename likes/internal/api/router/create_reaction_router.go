package router

import (
	"likes/config"
	"likes/internal/api/controller"

	"likes/internal/domain/repository"
	"likes/internal/domain/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCreateReactionRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	reactionRepository := repository.NewReactionRepository(db)
	CreateReactionController := controller.CreateReactionController{
		CreateReactionService: service.NewCreateReactionService(reactionRepository),
		Env:                   env,
	}
	gin.POST("/post/:post_id/reaction", CreateReactionController.Create)
}
