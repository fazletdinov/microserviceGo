package comment

import (
	"context"
	"posts/internal/domain/repository/comment"
	"posts/internal/models"
)

type createCommentService struct {
	commentRepository comment.CommentRepository
}

func NewCreateCommentService(commentRepository comment.CommentRepository) CreateCommentServcie {
	return &createCommentService{
		commentRepository: commentRepository,
	}
}

func (cs *createCommentService) CreateComment(ctx context.Context, comment *models.Comment) error {
	return cs.commentRepository.CreateComment(ctx, comment)
}
