package post

import (
	"context"
	"posts/internal/models"

	"github.com/google/uuid"
)

type PostGRPCRepository interface {
	Create(ctx context.Context, title string, content string, authorID uuid.UUID) (uuid.UUID, error)
	GetByIDPost(ctx context.Context, postID uuid.UUID) (*models.Post, error)
	GetPostByIDAuthorID(ctx context.Context, postID uuid.UUID, authorID uuid.UUID) (*models.Post, error)
	GetPosts(ctx context.Context, limit int, offset int) (*[]models.Post, error)
	UpdatePost(ctx context.Context, postID uuid.UUID, authorID uuid.UUID, title string, content string) error
	DeletePost(ctx context.Context, postID uuid.UUID, authorID uuid.UUID) error
	DeletePostsByAuthor(ctx context.Context, authorID uuid.UUID) error
}
