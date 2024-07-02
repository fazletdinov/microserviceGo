package service

import (
	"context"
	"posts/internal/domain/repository"
	"posts/internal/models"

	"github.com/google/uuid"
)

type deletePostService struct {
	postRepository repository.PostRepository
}

func NewDeletePostService(postRepository repository.PostRepository) DeletePostServcie {
	return &deletePostService{
		postRepository: postRepository,
	}
}

func (dps *deletePostService) GetByID(ctx context.Context, postID uuid.UUID) (*models.Post, error) {
	return dps.postRepository.GetByIDPost(ctx, postID)
}

func (dps *deletePostService) DeletePost(ctx context.Context, postID uuid.UUID) error {
	return dps.postRepository.DeletePost(ctx, postID)
}
