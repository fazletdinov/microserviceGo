package posts

import (
	"context"
	"posts/config"
	postsgrpc "posts/protogen/posts"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	commentService "posts/internal/domain/service/comment"
	postService "posts/internal/domain/service/post"
)

type GRPCController struct {
	postsgrpc.UnimplementedGatewayPostsServer
	PostGRPCService    postService.PostGRPCService
	CommentGRPCService commentService.CommentGRPCServcie
	Env                *config.Config
}

func (gc *GRPCController) CreatePost(
	ctx context.Context,
	postRequest *postsgrpc.CreatePostRequest,
) (*postsgrpc.CreatePostResponse, error) {
	if postRequest.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле title обязательно")
	}
	if postRequest.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле content обязательно")
	}
	if postRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	postID, err := gc.PostGRPCService.CreatePost(
		ctx,
		postRequest.GetTitle(),
		postRequest.GetContent(),
		uuid.MustParse(postRequest.GetAuthorId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &postsgrpc.CreatePostResponse{
		PostId: postID.String(),
	}, nil
}

func (gc *GRPCController) GetPostByID(
	ctx context.Context,
	postRequest *postsgrpc.GetPostRequest,
) (*postsgrpc.GetPostResponse, error) {
	if postRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}

	post, err := gc.PostGRPCService.GetByIDPost(
		ctx,
		uuid.MustParse(postRequest.GetPostId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	arrayComment := make([]*postsgrpc.Comment, 0, 10)
	for _, comment := range post.Comments {
		arrayComment = append(arrayComment, &postsgrpc.Comment{
			Text:      comment.Text,
			AuthorId:  comment.AuthorID.String(),
			CreatedAt: comment.CreatedAt.String(),
		})
	}
	return &postsgrpc.GetPostResponse{
		PostId:    post.ID.String(),
		Title:     post.Title,
		Content:   post.Content,
		AuthorId:  post.AuthorID.String(),
		CreatedAt: post.CreatedAt.String(),
		Comments:  arrayComment,
	}, nil
}

func (gc *GRPCController) GetPosts(
	ctx context.Context,
	postRequest *postsgrpc.GetPostsRequest,
) (*postsgrpc.GetPostsResponse, error) {
	posts, err := gc.PostGRPCService.GetPosts(
		ctx,
		postRequest.GetLimit(),
		postRequest.GetOffset(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	arrayPost := make([]*postsgrpc.Post, 0, postRequest.GetLimit())
	for _, post := range *posts {
		arrayComment := make([]*postsgrpc.Comment, 0, 10)
		for _, comment := range post.Comments {
			arrayComment = append(arrayComment, &postsgrpc.Comment{
				CommentId: comment.ID.String(),
				Text:      comment.Text,
				AuthorId:  comment.AuthorID.String(),
				CreatedAt: comment.CreatedAt.String(),
			})
		}
		arrayPost = append(arrayPost, &postsgrpc.Post{
			PostId:    post.ID.String(),
			Title:     post.Title,
			Content:   post.Content,
			AuthorId:  post.AuthorID.String(),
			CreatedAt: post.CreatedAt.String(),
			Comments:  arrayComment,
		})
	}
	return &postsgrpc.GetPostsResponse{
		Posts: arrayPost,
	}, nil
}

func (gc *GRPCController) UpdatePost(
	ctx context.Context,
	postRequest *postsgrpc.UpdatePostRequest,
) (*postsgrpc.UpdatePostResponse, error) {
	if postRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if postRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}
	if postRequest.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле title обязательно")
	}
	if postRequest.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле content обязательно")
	}
	err := gc.PostGRPCService.UpdatePost(
		ctx,
		uuid.MustParse(postRequest.GetPostId()),
		uuid.MustParse(postRequest.GetAuthorId()),
		postRequest.GetTitle(),
		postRequest.GetContent(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &postsgrpc.UpdatePostResponse{
		SuccessMessage: "Обновление прошло успешно",
	}, nil
}

func (gc *GRPCController) DeletePost(
	ctx context.Context,
	postRequest *postsgrpc.DeletePostRequest,
) (*postsgrpc.DeletePostResponse, error) {
	if postRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if postRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}
	err := gc.PostGRPCService.DeletePost(
		ctx,
		uuid.MustParse(postRequest.GetPostId()),
		uuid.MustParse(postRequest.GetAuthorId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &postsgrpc.DeletePostResponse{
		SuccessMessage: "Удаление прошло успешно",
	}, nil
}

// реализация контроллера для комментариев

func (gc *GRPCController) CreateComment(
	ctx context.Context,
	commentRequest *postsgrpc.CreateCommentRequest,
) (*postsgrpc.CreateCommentResponse, error) {
	if commentRequest.GetText() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле text обязательно")
	}
	if commentRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if commentRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	commentID, err := gc.CommentGRPCService.CreateComment(
		ctx,
		commentRequest.GetText(),
		uuid.MustParse(commentRequest.GetPostId()),
		uuid.MustParse(commentRequest.GetAuthorId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &postsgrpc.CreateCommentResponse{
		CommentId: commentID.String(),
	}, nil
}

func (gc *GRPCController) GetPostComments(
	ctx context.Context,
	commentRequest *postsgrpc.GetCommentsRequest,
) (*postsgrpc.GetCommentsResponse, error) {
	if commentRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}

	comments, err := gc.CommentGRPCService.GetPostComments(
		ctx,
		uuid.MustParse(commentRequest.GetPostId()),
		commentRequest.GetLimit(),
		commentRequest.GetOffset(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	arrayComment := make([]*postsgrpc.Comment, 0, commentRequest.GetLimit())
	for _, comment := range *comments {
		arrayComment = append(arrayComment, &postsgrpc.Comment{
			CommentId: comment.ID.String(),
			Text:      comment.Text,
			AuthorId:  comment.AuthorID.String(),
			CreatedAt: comment.CreatedAt.String(),
		})
	}
	return &postsgrpc.GetCommentsResponse{
		Comments: arrayComment,
	}, nil
}

func (gc *GRPCController) GetCommentByID(
	ctx context.Context,
	commentRequest *postsgrpc.GetCommentRequest,
) (*postsgrpc.GetCommentResponse, error) {
	if commentRequest.GetCommentId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле comment_id обязательно")
	}
	if commentRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if commentRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	comment, err := gc.CommentGRPCService.GetCommentByID(
		ctx,
		uuid.MustParse(commentRequest.GetCommentId()),
		uuid.MustParse(commentRequest.GetPostId()),
		uuid.MustParse(commentRequest.GetAuthorId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &postsgrpc.GetCommentResponse{
		Comment: &postsgrpc.Comment{
			CommentId: comment.ID.String(),
			Text:      comment.Text,
			AuthorId:  comment.AuthorID.String(),
			CreatedAt: comment.CreatedAt.String(),
		},
	}, nil
}

func (gc *GRPCController) UpdatePostComment(
	ctx context.Context,
	commentRequest *postsgrpc.UpdateCommentRequest,
) (*postsgrpc.UpdateCommentResponse, error) {
	if commentRequest.GetCommentId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле comment_id обязательно")
	}
	if commentRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if commentRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}
	if commentRequest.GetText() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле text обязательно")
	}

	err := gc.CommentGRPCService.UpdateComment(
		ctx,
		uuid.MustParse(commentRequest.GetCommentId()),
		uuid.MustParse(commentRequest.GetPostId()),
		uuid.MustParse(commentRequest.GetAuthorId()),
		commentRequest.GetText(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &postsgrpc.UpdateCommentResponse{
		SuccessMessage: "Обновление прошло успешно",
	}, nil
}

func (gc *GRPCController) DeletePostComment(
	ctx context.Context,
	commentRequest *postsgrpc.DeleteCommentRequest,
) (*postsgrpc.DeleteCommentResponse, error) {
	if commentRequest.GetCommentId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле comment_id обязательно")
	}
	if commentRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if commentRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	err := gc.CommentGRPCService.DeleteComment(
		ctx,
		uuid.MustParse(commentRequest.GetCommentId()),
		uuid.MustParse(commentRequest.GetPostId()),
		uuid.MustParse(commentRequest.GetAuthorId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &postsgrpc.DeleteCommentResponse{
		SuccessMessage: "Обновление прошло успешно",
	}, nil
}
