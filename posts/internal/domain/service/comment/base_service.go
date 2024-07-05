package comment

import (
	"context"
	"posts/internal/models"
	"posts/internal/schemas"

	"github.com/google/uuid"
)

type CreateCommentServcie interface {
	CreateComment(ctx context.Context, comment *models.Comment) error
}

type GetCommentServcie interface {
	GetByID(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID) (*models.Comment, error)
	GetComments(ctx context.Context, postID uuid.UUID, limit int, offset int) (*[]models.Comment, error)
}

type UpdateCommentService interface {
	GetByID(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID) (*models.Comment, error)
	UpdateComment(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID, comment *schemas.CommentUpdateRequest) error
}

type DeleteCommentService interface {
	GetByID(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID) (*models.Comment, error)
	DeleteComment(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID) error
}
