package service

import (
	"context"
	"posts/internal/models"
	"posts/internal/schemas"

	"github.com/google/uuid"
)

type CreatePostServcie interface {
	CreatePost(ctx context.Context, post *models.Post) error
}

type GetPostServcie interface {
	GetByID(ctx context.Context, postID uuid.UUID) (*models.Post, error)
	GetPosts(ctx context.Context, limit int, offset int) (*[]models.Post, error)
}

type UpdatePostServcie interface {
	GetByID(ctx context.Context, postID uuid.UUID) (*models.Post, error)
	UpdatePost(ctx context.Context, post *schemas.PostUpdateRequest) error
}

type DeletePostServcie interface {
	GetByID(ctx context.Context, postID uuid.UUID) (*models.Post, error)
	DeletePost(ctx context.Context, postID uuid.UUID) error
}
