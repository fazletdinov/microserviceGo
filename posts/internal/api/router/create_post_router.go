package router

import (
	"posts/config"
	"posts/internal/api/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"posts/internal/domain/repository"
	"posts/internal/domain/service"
)

func NewCreatePostRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(db)
	CreatePostController := controller.CreatePostController{
		CreatePosteService: service.NewCreatePostService(postRepository),
		Env:                env,
	}
	gin.POST("/create", CreatePostController.Create)
}
