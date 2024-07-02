package service

import (
	"context"
	"posts/internal/domain/repository"
	"posts/internal/models"

	"github.com/google/uuid"
)

type getPostService struct {
	postRepository repository.PostRepository
}

func NewGetPostService(postRepository repository.PostRepository) GetPostServcie {
	return &getPostService{
		postRepository: postRepository,
	}
}

func (pc *getPostService) GetByID(ctx context.Context, postID uuid.UUID) (*models.Post, error) {
	return pc.postRepository.GetByIDPost(ctx, postID)
}

func (pc *getPostService) GetPosts(ctx context.Context, limit int, offset int) (*[]models.Post, error) {
	return pc.postRepository.GetPosts(ctx, limit, offset)
}
