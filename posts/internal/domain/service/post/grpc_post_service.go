package post

import (
	"context"
	"posts/internal/domain/repository/post"
	"posts/internal/models"

	"github.com/google/uuid"
)

type postService struct {
	postRepository post.PostRepository
}

func NewPostService(postRepository post.PostRepository) PostService {
	return &postService{
		postRepository: postRepository,
	}
}

func (ps *postService) CreatePost(
	ctx context.Context,
	title string,
	content string,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	return ps.postRepository.Create(ctx, title, content, authorID)
}

func (ps *postService) GetByIDPost(
	ctx context.Context,
	postID uuid.UUID,
) (*models.Post, error) {
	return ps.postRepository.GetByIDPost(ctx, postID)
}

func (ps *postService) GetPostByIDAuthorID(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*models.Post, error) {
	return ps.postRepository.GetPostByIDAuthorID(ctx, postID, authorID)
}

func (ps *postService) GetPosts(
	ctx context.Context,
	limit uint64,
	offset uint64,
) (*[]models.Post, error) {
	return ps.postRepository.GetPosts(ctx, int(limit), int(offset))
}

func (ps *postService) UpdatePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
	title string,
	content string,
) error {
	return ps.postRepository.UpdatePost(ctx, postID, authorID, title, content)
}

func (ps *postService) DeletePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	return ps.postRepository.DeletePost(ctx, postID, authorID)
}

func (ps *postService) DeletePostsByAuthor(
	ctx context.Context,
	authorID uuid.UUID,
) error {
	return ps.postRepository.DeletePostsByAuthor(ctx, authorID)
}
