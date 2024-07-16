package post

import (
	"context"
	"posts/internal/domain/repository/post"
	"posts/internal/models"

	"github.com/google/uuid"
)

type postGRPCService struct {
	postGRPCRepository post.PostGRPCRepository
}

func NewPostGRPCService(postGRPCRepository post.PostGRPCRepository) PostGRPCService {
	return &postGRPCService{
		postGRPCRepository: postGRPCRepository,
	}
}

func (ps *postGRPCService) CreatePost(
	ctx context.Context,
	title string,
	content string,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	return ps.postGRPCRepository.Create(ctx, title, content, authorID)
}

func (ps *postGRPCService) GetByIDPost(
	ctx context.Context,
	postID uuid.UUID,
) (*models.Post, error) {
	return ps.postGRPCRepository.GetByIDPost(ctx, postID)
}

func (ps *postGRPCService) GetPosts(
	ctx context.Context,
	limit uint64,
	offset uint64,
) (*[]models.Post, error) {
	return ps.postGRPCRepository.GetPosts(ctx, int(limit), int(offset))
}

func (ps *postGRPCService) UpdatePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
	title string,
	content string,
) error {
	return ps.postGRPCRepository.UpdatePost(ctx, postID, authorID, title, content)
}

func (ps *postGRPCService) DeletePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	return ps.postGRPCRepository.DeletePost(ctx, postID, authorID)
}
