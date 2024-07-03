package router

import (
	"likes/config"
	"likes/internal/api/controller"
	"likes/internal/domain/repository"
	"likes/internal/domain/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewGetReactionRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	reactionRepository := repository.NewReactionRepository(db)
	GetReactionController := controller.GetReactionController{
		GetReactionService: service.NewGetReactionService(reactionRepository),
		Env:                env,
	}
	gin.GET("/post/:post_id/reactions", GetReactionController.Fetchs)
}
