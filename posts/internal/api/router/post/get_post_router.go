package post

import (
	"posts/config"
	postController "posts/internal/api/controller/post"
	postrepository "posts/internal/domain/repository/post"
	postService "posts/internal/domain/service/post"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewGetPostRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	postRepository := postrepository.NewPostRepository(db)
	GetPostController := postController.GetPostController{
		GetPostService: postService.NewGetPostService(postRepository),
		Env:            env,
	}
	gin.GET("/post/:post_id", GetPostController.Fetch)
	gin.GET("/posts", GetPostController.Fetchs)
}
