package post

import (
	"context"
	"posts/internal/domain/repository/post"
	"posts/internal/models"
)

type postCreateService struct {
	postRepository post.PostRepository
}

func NewCreatePostService(postRepository post.PostRepository) CreatePostServcie {
	return &postCreateService{
		postRepository: postRepository,
	}
}

func (pc *postCreateService) CreatePost(ctx context.Context, post *models.Post) error {
	return pc.postRepository.Create(ctx, post)
}
