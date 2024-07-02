package post

import (
	"context"
	"posts/internal/domain/repository/post"
	"posts/internal/models"
	"posts/internal/schemas"

	"github.com/google/uuid"
)

type updatePostService struct {
	postRepository post.PostRepository
}

func NewUpdatePostService(postRepository post.PostRepository) UpdatePostServcie {
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
