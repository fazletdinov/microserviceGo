package post

import (
	"context"
	"posts/internal/models"
	"posts/internal/schemas"

	"github.com/google/uuid"
)

type PostRepository interface {
	Create(ctx context.Context, post *models.Post) error
	GetByIDPost(ctx context.Context, postID uuid.UUID) (*models.Post, error)
	GetPosts(ctx context.Context, limit int, offset int) (*[]models.Post, error)
	UpdatePost(ctx context.Context, post *schemas.PostUpdateRequest) error
	DeletePost(ctx context.Context, postID uuid.UUID) error
}
