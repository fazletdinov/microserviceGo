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

func NewUpdateCommentRouter(env *config.Config, db *gorm.DB, gin *gin.RouterGroup) {
	commentRepository := commentRepository.NewCommentRepository(db)
	postRepository := postRepository.NewPostRepository(db)
	UpdateCommentController := commentController.UpdateCommentController{
		UpdateCommentService: commentService.NewUpdateCommentService(commentRepository),
		GetPostService:       postService.NewGetPostService(postRepository),
		Env:                  env,
	}
	gin.PUT("/post/:post_id/comment/:comment_id", UpdateCommentController.Update)
}
