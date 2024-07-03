package router

import (
	"likes/config"
	"likes/internal/api/controller"
	"likes/internal/domain/repository"
	"likes/internal/domain/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewDeleteReactionRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	reactionRepository := repository.NewReactionRepository(db)
	DeleteReactionController := controller.DeleteReactionController{
		DeleteReactionService: service.NewDeleteReactionService(reactionRepository),
		Env:                   env,
	}

	gin.DELETE("/post/:post_id/reaction", DeleteReactionController.Delete)
}
