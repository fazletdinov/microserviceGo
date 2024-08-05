package post

import (
	"context"
	"posts/internal/domain/repository/post"
	"posts/internal/dto"

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
) (*dto.PostResponse, error) {
	postResponse, err := ps.postRepository.GetByIDPost(ctx, postID)
	if err != nil {
		return nil, err
	}
	return &dto.PostResponse{
		ID:        postResponse.ID,
		Title:     postResponse.Title,
		Content:   postResponse.Content,
		AuthorID:  postResponse.AuthorID,
		CreatedAt: postResponse.CreatedAt,
		Comments:  &postResponse.Comments,
	}, nil
}

func (ps *postService) GetPostByIDAuthorID(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*dto.PostResponse, error) {
	postResponse, err := ps.postRepository.GetPostByIDAuthorID(ctx, postID, authorID)
	if err != nil {
		return nil, err
	}
	return &dto.PostResponse{
		ID:        postResponse.ID,
		Title:     postResponse.Title,
		Content:   postResponse.Content,
		AuthorID:  postResponse.AuthorID,
		CreatedAt: postResponse.CreatedAt,
		Comments:  &postResponse.Comments,
	}, nil
}

func (ps *postService) GetPosts(
	ctx context.Context,
	limit uint64,
	offset uint64,
) (*[]dto.PostResponse, error) {
	postsResponse, err := ps.postRepository.GetPosts(ctx, int(limit), int(offset))
	if err != nil {
		return nil, err
	}
	posts := make([]dto.PostResponse, 0, limit)
	for _, post := range *postsResponse {
		posts = append(posts, dto.PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			AuthorID:  post.AuthorID,
			CreatedAt: post.CreatedAt,
			Comments:  &post.Comments,
		})
	}

	return &posts, nil
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
