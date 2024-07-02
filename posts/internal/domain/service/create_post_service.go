package service

import (
	"context"
	"posts/internal/domain/repository"
	"posts/internal/models"
)

type postCreateService struct {
	postRepository repository.PostRepository
}

func NewCreatePostService(postRepository repository.PostRepository) CreatePostServcie {
	return &postCreateService{
		postRepository: postRepository,
	}
}

func (pc *postCreateService) CreatePost(ctx context.Context, post *models.Post) error {
	return pc.postRepository.Create(ctx, post)
}
