package comment

import (
	"context"
	"posts/internal/domain/repository/comment"
	"posts/internal/models"

	"github.com/google/uuid"
)

type deleteCommentService struct {
	commentRepository comment.CommentRepository
}

func NewDeleteCommentService(commentRepository comment.CommentRepository) DeleteCommentService {
	return &deleteCommentService{
		commentRepository: commentRepository,
	}
}

func (ucs *deleteCommentService) GetByID(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID) (*models.Comment, error) {
	return ucs.commentRepository.GetByIDComment(ctx, postID, commentID, authorID)
}

func (ucs *deleteCommentService) DeleteComment(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID) error {
	return ucs.commentRepository.DeleteComment(ctx, postID, commentID, authorID)
}
