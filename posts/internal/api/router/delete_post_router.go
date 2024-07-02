package router

import (
	"posts/config"
	"posts/internal/api/controller"
	"posts/internal/domain/repository"
	"posts/internal/domain/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewDeletePostRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(db)
	DeletePostController := controller.DeletePostController{
		DeletePostService: service.NewDeletePostService(postRepository),
		Env:               env,
	}

	gin.DELETE("/delete/:id", DeletePostController.Delete)
}
