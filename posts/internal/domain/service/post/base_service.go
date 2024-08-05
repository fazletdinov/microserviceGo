package post

import (
	"context"
	"posts/internal/dto"

	"github.com/google/uuid"
)

type PostService interface {
	CreatePost(ctx context.Context, title string, content string, authorID uuid.UUID) (uuid.UUID, error)
	GetByIDPost(ctx context.Context, postID uuid.UUID) (*dto.PostResponse, error)
	GetPostByIDAuthorID(ctx context.Context, postID uuid.UUID, authorID uuid.UUID) (*dto.PostResponse, error)
	GetPosts(ctx context.Context, limit uint64, offset uint64) (*[]dto.PostResponse, error)
	UpdatePost(ctx context.Context, postID uuid.UUID, authorID uuid.UUID, title string, content string) error
	DeletePost(ctx context.Context, postID uuid.UUID, authorID uuid.UUID) error
	DeletePostsByAuthor(ctx context.Context, authorID uuid.UUID) error
}
