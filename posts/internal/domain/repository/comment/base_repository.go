package comment

import (
	"context"
	"posts/internal/models"

	"github.com/google/uuid"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, text string, postID uuid.UUID, authorID uuid.UUID) (uuid.UUID, error)
	GetByIDComment(ctx context.Context, commentID uuid.UUID, postID uuid.UUID, authorID uuid.UUID) (*models.Comment, error)
	GetComments(ctx context.Context, postID uuid.UUID, limit int, offset int) (*[]models.Comment, error)
	UpdateComment(ctx context.Context, commentID uuid.UUID, postID uuid.UUID, authorID uuid.UUID, text string) error
	DeleteComment(ctx context.Context, commentID uuid.UUID, postID uuid.UUID, authorID uuid.UUID) error
}
