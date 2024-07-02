package comment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"posts/config"
	commentController "posts/internal/api/controller/comment"
	commentRepository "posts/internal/domain/repository/comment"
	postRepository "posts/internal/domain/repository/post"
	commentService "posts/internal/domain/service/comment"
	postService "posts/internal/domain/service/post"
)

func NewCreateCommentRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	commentRepository := commentRepository.NewCommentRepository(db)
	postRepository := postRepository.NewPostRepository(db)
	CreateCommentController := commentController.CreateCommentController{
		CreateCommentService: commentService.NewCreateCommentService(commentRepository),
		GetPostService:       postService.NewGetPostService(postRepository),
		Env:                  env,
	}
	gin.POST("/post/:post_id/comment", CreateCommentController.Create)
}
