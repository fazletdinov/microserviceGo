package post

import (
	"posts/config"
	postController "posts/internal/api/controller/post"
	postrepository "posts/internal/domain/repository/post"
	postService "posts/internal/domain/service/post"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewDeletePostRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	postRepository := postrepository.NewPostRepository(db)
	DeletePostController := postController.DeletePostController{
		DeletePostService: postService.NewDeletePostService(postRepository),
		Env:               env,
	}

	gin.DELETE("/post/:post_id", DeletePostController.Delete)
}
