package posts

import (
	"api-grpc-gateway/config"
	postsgrpc "api-grpc-gateway/protogen/golang/posts"
	"context"

	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClientPosts struct {
	posts postsgrpc.GatewayPostsClient
	env   *config.Config
}

func NewGRPCClientPosts(
	addrs string,
	env *config.Config,
) (*GRPCClientPosts, error) {
	cc, err := grpc.NewClient(
		addrs,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}
	return &GRPCClientPosts{
		posts: postsgrpc.NewGatewayPostsClient(cc),
		env:   env,
	}, nil
}

func (gc *GRPCClientPosts) CreatePost(
	ctx context.Context,
	title string,
	content string,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"CreatePost",
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
		oteltrace.WithAttributes(attribute.String("Title", title)),
		oteltrace.WithAttributes(attribute.String("Content", content)),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для создания post")
	defer span.End()
	postID, err := gc.posts.CreatePost(traceCtx, &postsgrpc.CreatePostRequest{
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
) (*postsgrpc.PostResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"GetPostByID",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для получения post по ID")
	defer span.End()
	postResponse, err := gc.posts.GetPostByID(traceCtx, &postsgrpc.GetPostRequest{
		PostId: postID.String(),
	})

	if err != nil {
		return nil, err
	}

	return postResponse, nil
}

func (gc *GRPCClientPosts) GetPostByIDAuthorID(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*postsgrpc.PostResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"GetPostByIDAuthorID",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для получения post по ID и AuthorID")
	defer span.End()
	postResponse, err := gc.posts.GetPostByIDAuthorID(traceCtx, &postsgrpc.GetPostByIDAuthorIDRequest{
		PostId:   postID.String(),
		AuthorId: authorID.String(),
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
) ([]*postsgrpc.PostResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"GetPosts",
		oteltrace.WithAttributes(attribute.Int64("Limit", int64(limit))),
		oteltrace.WithAttributes(attribute.Int64("Offset", int64(offset))),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для получения всех posts")
	defer span.End()
	postResponse, err := gc.posts.GetPosts(traceCtx, &postsgrpc.GetPostsRequest{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		return nil, err
	}

	return postResponse.Posts, nil
}

func (gc *GRPCClientPosts) UpdatePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
	title string,
	content string,
) (*postsgrpc.UpdatePostResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"UpdatePost",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
		oteltrace.WithAttributes(attribute.String("Title", title)),
		oteltrace.WithAttributes(attribute.String("Content", content)),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для обновления post")
	defer span.End()
	postResponse, err := gc.posts.UpdatePost(traceCtx, &postsgrpc.UpdatePostRequest{
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
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"DeletePost",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для удаления post")
	defer span.End()
	postResponse, err := gc.posts.DeletePost(traceCtx, &postsgrpc.DeletePostRequest{
		PostId:   postID.String(),
		AuthorId: authorID.String(),
	})

	if err != nil {
		return nil, err
	}

	return postResponse, nil
}

func (gc *GRPCClientPosts) DeletePostsByAuthor(
	ctx context.Context,
	authorID uuid.UUID,
) (*postsgrpc.DeletePostsByAuthorResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"DeletePostsByAuthor",
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для удаления posts пользователя")
	defer span.End()
	postResponse, err := gc.posts.DeletePostsByAuthor(traceCtx, &postsgrpc.DeletePostsByAuthorRequest{
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
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"CreateComment",
		oteltrace.WithAttributes(attribute.String("Text", text)),
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для создания comment")
	defer span.End()
	commentResponse, err := gc.posts.CreateComment(traceCtx, &postsgrpc.CreateCommentRequest{
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
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"GetCommentByID",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
		oteltrace.WithAttributes(attribute.String("CommentID", commentID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для получения comment по ID")
	defer span.End()
	postResponse, err := gc.posts.GetCommentByID(traceCtx, &postsgrpc.GetCommentRequest{
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
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"GetPostComments",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.Int64("Limit", int64(limit))),
		oteltrace.WithAttributes(attribute.Int64("Offset", int64(offset))),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для получения comments на post")
	defer span.End()
	commentResponse, err := gc.posts.GetPostComments(traceCtx, &postsgrpc.GetCommentsRequest{
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
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"UpdateComment",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
		oteltrace.WithAttributes(attribute.String("CommentID", commentID.String())),
		oteltrace.WithAttributes(attribute.String("Text", text)),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для обновления comment")
	defer span.End()
	commentResponse, err := gc.posts.UpdatePostComment(traceCtx, &postsgrpc.UpdateCommentRequest{
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
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"DeleteComment",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
		oteltrace.WithAttributes(attribute.String("CommentID", commentID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис posts для удаления comment")
	defer span.End()
	commentResponse, err := gc.posts.DeletePostComment(traceCtx, &postsgrpc.DeleteCommentRequest{
		CommentId: commentID.String(),
		PostId:    postID.String(),
		AuthorId:  authorID.String(),
	})

	if err != nil {
		return nil, err
	}

	return commentResponse, nil
}
