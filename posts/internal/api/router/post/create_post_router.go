package post

import (
	"posts/config"
	postController "posts/internal/api/controller/post"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	postRepository "posts/internal/domain/repository/post"
	postService "posts/internal/domain/service/post"
)

func NewCreatePostRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	postRepository := postRepository.NewPostRepository(db)
	CreatePostController := postController.CreatePostController{
		CreatePosteService: postService.NewCreatePostService(postRepository),
		Env:                env,
	}
	gin.POST("/post", CreatePostController.Create)
}
