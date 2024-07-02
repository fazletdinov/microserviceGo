package service

import (
	"context"
	"posts/internal/domain/repository"
	"posts/internal/models"
	"posts/internal/schemas"

	"github.com/google/uuid"
)

type updatePostService struct {
	postRepository repository.PostRepository
}

func NewUpdatePostService(postRepository repository.PostRepository) UpdatePostServcie {
	return &updatePostService{
		postRepository: postRepository,
	}
}

func (ups *updatePostService) GetByID(ctx context.Context, postID uuid.UUID) (*models.Post, error) {
	return ups.postRepository.GetByIDPost(ctx, postID)
}

func (ups *updatePostService) UpdatePost(ctx context.Context, post *schemas.PostUpdateRequest) error {
	return ups.postRepository.UpdatePost(ctx, post)
}
