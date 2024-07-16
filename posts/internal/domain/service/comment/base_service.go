package comment

import (
	"context"
	"posts/internal/models"

	"github.com/google/uuid"
)

type CommentGRPCServcie interface {
	CreateComment(ctx context.Context, text string, postID uuid.UUID, authorID uuid.UUID) (uuid.UUID, error)
	GetPostComments(ctx context.Context, postID uuid.UUID, limit uint64, offset uint64) (*[]models.Comment, error)
	GetCommentByID(ctx context.Context, commentID uuid.UUID, postID uuid.UUID, authorID uuid.UUID) (*models.Comment, error)
	UpdateComment(ctx context.Context, commentID uuid.UUID, postID uuid.UUID, authorID uuid.UUID, text string) error
	DeleteComment(ctx context.Context, commentID uuid.UUID, postID uuid.UUID, authorID uuid.UUID) error
}
