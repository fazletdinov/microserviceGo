package comment

import (
	"context"
	"posts/internal/domain/repository/comment"
	"posts/internal/models"

	"github.com/google/uuid"
)

type getCommentService struct {
	commentRepository comment.CommentRepository
}

func NewGetCommentService(commentRepository comment.CommentRepository) GetCommentServcie {
	return &getCommentService{
		commentRepository: commentRepository,
	}
}

func (gcs *getCommentService) GetByID(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID) (*models.Comment, error) {
	return gcs.commentRepository.GetByIDComment(ctx, postID, commentID, authorID)
}

func (gcs *getCommentService) GetComments(ctx context.Context, postID uuid.UUID, limit int, offset int) (*[]models.Comment, error) {
	return gcs.commentRepository.GetComments(ctx, postID, limit, offset)
}
