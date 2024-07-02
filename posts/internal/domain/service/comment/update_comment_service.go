package comment

import (
	"context"
	"posts/internal/domain/repository/comment"
	"posts/internal/models"
	"posts/internal/schemas"

	"github.com/google/uuid"
)

type updateCommentService struct {
	commentRepository comment.CommentRepository
}

func NewUpdateCommentService(commentRepository comment.CommentRepository) UpdateCommentService {
	return &updateCommentService{
		commentRepository: commentRepository,
	}
}

func (ucs *updateCommentService) GetByID(ctx context.Context, postID uuid.UUID, commentID uuid.UUID) (*models.Comment, error) {
	return ucs.commentRepository.GetByIDComment(ctx, postID, commentID)
}

func (ucs *updateCommentService) UpdateComment(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, comment *schemas.CommentUpdateRequest) error {
	return ucs.commentRepository.UpdateComment(ctx, postID, commentID, comment)
}
