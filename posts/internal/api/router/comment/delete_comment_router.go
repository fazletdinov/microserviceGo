package comment

import (
	"posts/config"
	commentController "posts/internal/api/controller/comment"
	commentRepository "posts/internal/domain/repository/comment"
	postRepository "posts/internal/domain/repository/post"
	commentService "posts/internal/domain/service/comment"
	postService "posts/internal/domain/service/post"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewDeleteCommentRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	commentRepository := commentRepository.NewCommentRepository(db)
	postRepository := postRepository.NewPostRepository(db)
	DeleteCommentController := commentController.DeleteCommentController{
		DeleteCommentService: commentService.NewDeleteCommentService(commentRepository),
		GetPostService:       postService.NewGetPostService(postRepository),
		Env:                  env,
	}
	gin.DELETE("/post/:post_id/comment/:comment_id", DeleteCommentController.Delete)
}
