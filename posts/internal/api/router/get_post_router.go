package router

import (
	"posts/config"
	"posts/internal/api/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"posts/internal/domain/repository"
	"posts/internal/domain/service"
)

func NewGetPostRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(db)
	GetPostController := controller.GetPostController{
		GetPostService: service.NewGetPostService(postRepository),
		Env:            env,
	}
	gin.GET("/post/:id", GetPostController.Fetch)
	gin.GET("/posts", GetPostController.Fetchs)
}
