package posts

import (
	postsgrpc "api-grpc-gateway/protogen/golang/posts"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClientPosts struct {
	posts postsgrpc.GatewayPostsClient
}

func NewGRPCClientPosts(
	addrs string,
) (*GRPCClientPosts, error) {
	cc, err := grpc.NewClient(addrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &GRPCClientPosts{posts: postsgrpc.NewGatewayPostsClient(cc)}, nil
}

func (gc *GRPCClientPosts) CreatePost(
	ctx context.Context,
	title string,
	content string,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	postID, err := gc.posts.CreatePost(ctx, &postsgrpc.CreatePostRequest{
		Title:    title,
		Content:  content,
		AuthorId: authorID.String(),
	})

	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(postID.PostId), nil
}

func (gc *GRPCClientPosts) GetPostByID(
	ctx context.Context,
	postID uuid.UUID,
) (*postsgrpc.GetPostResponse, error) {
	postResponse, err := gc.posts.GetPostByID(ctx, &postsgrpc.GetPostRequest{
		PostId: postID.String(),
	})

	if err != nil {
		return nil, err
	}

	return postResponse, nil
}

func (gc *GRPCClientPosts) GetPosts(
	ctx context.Context,
	limit uint64,
	offset uint64,
) (*postsgrpc.GetPostsResponse, error) {
	postResponse, err := gc.posts.GetPosts(ctx, &postsgrpc.GetPostsRequest{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		return nil, err
	}

	return postResponse, nil
}

func (gc *GRPCClientPosts) UpdatePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
	title string,
	content string,
) (*postsgrpc.UpdatePostResponse, error) {
	postResponse, err := gc.posts.UpdatePost(ctx, &postsgrpc.UpdatePostRequest{
		PostId:   postID.String(),
		AuthorId: authorID.String(),
		Title:    title,
		Content:  content,
	})

	if err != nil {
		return nil, err
	}

	return postResponse, nil
}

func (gc *GRPCClientPosts) DeletePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*postsgrpc.DeletePostResponse, error) {
	postResponse, err := gc.posts.DeletePost(ctx, &postsgrpc.DeletePostRequest{
		PostId:   postID.String(),
		AuthorId: authorID.String(),
	})

	if err != nil {
		return nil, err
	}

	return postResponse, nil
}

func (gc *GRPCClientPosts) CreateComment(
	ctx context.Context,
	text string,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*postsgrpc.CreateCommentResponse, error) {
	commentResponse, err := gc.posts.CreateComment(ctx, &postsgrpc.CreateCommentRequest{
		Text:     text,
		PostId:   postID.String(),
		AuthorId: authorID.String(),
	})

	if err != nil {
		return nil, err
	}

	return commentResponse, nil
}

func (gc *GRPCClientPosts) GetCommentByID(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*postsgrpc.GetCommentResponse, error) {
	postResponse, err := gc.posts.GetCommentByID(ctx, &postsgrpc.GetCommentRequest{
		CommentId: commentID.String(),
		PostId:    postID.String(),
		AuthorId:  authorID.String(),
	})

	if err != nil {
		return nil, err
	}

	return postResponse, nil
}

func (gc *GRPCClientPosts) GetPostComments(
	ctx context.Context,
	postID uuid.UUID,
	limit uint64,
	offset uint64,
) (*postsgrpc.GetCommentsResponse, error) {
	commentResponse, err := gc.posts.GetPostComments(ctx, &postsgrpc.GetCommentsRequest{
		PostId: postID.String(),
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		return nil, err
	}

	return commentResponse, nil
}

func (gc *GRPCClientPosts) UpdateComment(
	ctx context.Context,
	commentID uuid.UUID,
	text string,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*postsgrpc.UpdateCommentResponse, error) {
	commentResponse, err := gc.posts.UpdatePostComment(ctx, &postsgrpc.UpdateCommentRequest{
		CommentId: commentID.String(),
		Text:      text,
		PostId:    postID.String(),
		AuthorId:  authorID.String(),
	})

	if err != nil {
		return nil, err
	}

	return commentResponse, nil
}

func (gc *GRPCClientPosts) DeleteComment(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*postsgrpc.DeleteCommentResponse, error) {
	commentResponse, err := gc.posts.DeletePostComment(ctx, &postsgrpc.DeleteCommentRequest{
		CommentId: commentID.String(),
		PostId:    postID.String(),
		AuthorId:  authorID.String(),
	})

	if err != nil {
		return nil, err
	}

	return commentResponse, nil
}
