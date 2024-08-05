package comment

import (
	"context"
	"posts/internal/dto"

	"github.com/google/uuid"
)

type CommentServcie interface {
	CreateComment(ctx context.Context, text string, postID uuid.UUID, authorID uuid.UUID) (uuid.UUID, error)
	GetPostComments(ctx context.Context, postID uuid.UUID, limit uint64, offset uint64) (*[]dto.CommentResponse, error)
	GetCommentByID(ctx context.Context, commentID uuid.UUID, postID uuid.UUID, authorID uuid.UUID) (*dto.CommentResponse, error)
	UpdateComment(ctx context.Context, commentID uuid.UUID, postID uuid.UUID, authorID uuid.UUID, text string) error
	DeleteComment(ctx context.Context, commentID uuid.UUID, postID uuid.UUID, authorID uuid.UUID) error
}
