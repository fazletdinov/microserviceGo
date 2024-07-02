package router

import (
	"posts/config"
	"posts/internal/api/controller"
	"posts/internal/domain/repository"
	"posts/internal/domain/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUpdatePostRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(db)
	UpdatePostController := controller.UpdatePostController{
		UpdatePostService: service.NewUpdatePostService(postRepository),
		Env:               env,
	}
	gin.PUT("/update/:id", UpdatePostController.Update)
}
