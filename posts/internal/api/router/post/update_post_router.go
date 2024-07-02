package post

import (
	"posts/config"
	postController "posts/internal/api/controller/post"
	postrepository "posts/internal/domain/repository/post"
	postService "posts/internal/domain/service/post"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUpdatePostRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	postRepository := postrepository.NewPostRepository(db)
	UpdatePostController := postController.UpdatePostController{
		UpdatePostService: postService.NewUpdatePostService(postRepository),
		Env:               env,
	}
	gin.PUT("/post/:post_id", UpdatePostController.Update)
}
