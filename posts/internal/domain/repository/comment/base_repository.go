package comment

import (
	"context"
	"github.com/google/uuid"
	"posts/internal/models"
	"posts/internal/schemas"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *models.Comment) error
	GetByIDComment(ctx context.Context, postID uuid.UUID, commentID uuid.UUID) (*models.Comment, error)
	GetComments(ctx context.Context, postID uuid.UUID, limit int, offset int) (*[]models.Comment, error)
	UpdateComment(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, comment *schemas.CommentUpdateRequest) error
	DeleteComment(ctx context.Context, postID uuid.UUID, commentID uuid.UUID) error
}
