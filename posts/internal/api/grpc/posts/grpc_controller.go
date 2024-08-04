package posts

import (
	"context"
	"posts/config"
	postsgrpc "posts/protogen/posts"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
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
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if postRequest.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле title обязательно")
	}
	if postRequest.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле content обязательно")
	}
	if postRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"CreatePost",
		oteltrace.WithAttributes(attribute.String("AuthorID", postRequest.AuthorId)),
		oteltrace.WithAttributes(attribute.String("Title", postRequest.Title)),
		oteltrace.WithAttributes(attribute.String("Content", postRequest.Content)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для создания Post")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"CreatePost_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	postID, err := gc.PostGRPCService.CreatePost(
		traceCtx,
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
) (*postsgrpc.PostResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if postRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"GetPostByID",
		oteltrace.WithAttributes(attribute.String("PostID", postRequest.PostId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для получения Post по ID")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"GetPostByID_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	post, err := gc.PostGRPCService.GetByIDPost(
		traceCtx,
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
	return &postsgrpc.PostResponse{
		PostId:    post.ID.String(),
		Title:     post.Title,
		Content:   post.Content,
		AuthorId:  post.AuthorID.String(),
		CreatedAt: post.CreatedAt.String(),
		Comments:  arrayComment,
	}, nil
}

func (gc *GRPCController) GetPostByIDAuthorID(
	ctx context.Context,
	postRequest *postsgrpc.GetPostByIDAuthorIDRequest,
) (*postsgrpc.PostResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if postRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if postRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"GetPostByIDAuthorID",
		oteltrace.WithAttributes(attribute.String("PostID", postRequest.PostId)),
		oteltrace.WithAttributes(attribute.String("AuthorID", postRequest.AuthorId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для получения Post по ID и AuthorID")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"GetPostByIDAuthorID_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	post, err := gc.PostGRPCService.GetPostByIDAuthorID(
		traceCtx,
		uuid.MustParse(postRequest.GetPostId()),
		uuid.MustParse(postRequest.GetAuthorId()),
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
	return &postsgrpc.PostResponse{
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
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	traceCtx, span := tracer.Start(
		ctx,
		"GetPosts",
		oteltrace.WithAttributes(attribute.Int64("Limit", int64(postRequest.Limit))),
		oteltrace.WithAttributes(attribute.Int64("Offset", int64(postRequest.Offset))),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для получения Posts")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"GetPosts_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	posts, err := gc.PostGRPCService.GetPosts(
		traceCtx,
		postRequest.GetLimit(),
		postRequest.GetOffset(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	arrayPost := make([]*postsgrpc.PostResponse, 0, postRequest.GetLimit())
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
		arrayPost = append(arrayPost, &postsgrpc.PostResponse{
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
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

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

	traceCtx, span := tracer.Start(
		ctx,
		"UpdatePost",
		oteltrace.WithAttributes(attribute.String("PostID", postRequest.PostId)),
		oteltrace.WithAttributes(attribute.String("AuthorID", postRequest.AuthorId)),
		oteltrace.WithAttributes(attribute.String("Title", postRequest.Title)),
		oteltrace.WithAttributes(attribute.String("Content", postRequest.Content)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для обновления Post")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"UpdatePost_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	err := gc.PostGRPCService.UpdatePost(
		traceCtx,
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
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if postRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if postRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"DeletePost",
		oteltrace.WithAttributes(attribute.String("PostID", postRequest.PostId)),
		oteltrace.WithAttributes(attribute.String("AuthorID", postRequest.AuthorId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для удаления Post")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"DeletePost_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	err := gc.PostGRPCService.DeletePost(
		traceCtx,
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

func (gc *GRPCController) DeletePostsByAuthor(
	ctx context.Context,
	postRequest *postsgrpc.DeletePostsByAuthorRequest,
) (*postsgrpc.DeletePostsByAuthorResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if postRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"DeletePostsByAuthor",
		oteltrace.WithAttributes(attribute.String("AuthorID", postRequest.AuthorId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для удаления Posts автора")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"DeletePostsByAuthor_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	err := gc.PostGRPCService.DeletePostsByAuthor(
		traceCtx,
		uuid.MustParse(postRequest.GetAuthorId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &postsgrpc.DeletePostsByAuthorResponse{
		SuccessMessage: "Удаление прошло успешно",
	}, nil
}

// реализация контроллера для комментариев

func (gc *GRPCController) CreateComment(
	ctx context.Context,
	commentRequest *postsgrpc.CreateCommentRequest,
) (*postsgrpc.CreateCommentResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if commentRequest.GetText() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле text обязательно")
	}
	if commentRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if commentRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"CreateComment",
		oteltrace.WithAttributes(attribute.String("PostID", commentRequest.PostId)),
		oteltrace.WithAttributes(attribute.String("AuthorID", commentRequest.AuthorId)),
		oteltrace.WithAttributes(attribute.String("Text", commentRequest.Text)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для создания Comment")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"CreateComment_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	commentID, err := gc.CommentGRPCService.CreateComment(
		traceCtx,
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
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if commentRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"GetPostComments",
		oteltrace.WithAttributes(attribute.String("PostID", commentRequest.PostId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для получения Comments на Post")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"GetPostComments_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	comments, err := gc.CommentGRPCService.GetPostComments(
		traceCtx,
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
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if commentRequest.GetCommentId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле comment_id обязательно")
	}
	if commentRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if commentRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"GetCommentByID",
		oteltrace.WithAttributes(attribute.String("CommentID", commentRequest.CommentId)),
		oteltrace.WithAttributes(attribute.String("PostID", commentRequest.PostId)),
		oteltrace.WithAttributes(attribute.String("AuthorID", commentRequest.AuthorId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для получения Comment по ID")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"GetCommentByID_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	comment, err := gc.CommentGRPCService.GetCommentByID(
		traceCtx,
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
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

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

	traceCtx, span := tracer.Start(
		ctx,
		"UpdatePostComment",
		oteltrace.WithAttributes(attribute.String("CommentID", commentRequest.CommentId)),
		oteltrace.WithAttributes(attribute.String("PostID", commentRequest.PostId)),
		oteltrace.WithAttributes(attribute.String("AuthorID", commentRequest.AuthorId)),
		oteltrace.WithAttributes(attribute.String("Text", commentRequest.Text)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для обновления Comment по ID")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"UpdatePostComment_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	err := gc.CommentGRPCService.UpdateComment(
		traceCtx,
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
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if commentRequest.GetCommentId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле comment_id обязательно")
	}
	if commentRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if commentRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"DeletePostComment",
		oteltrace.WithAttributes(attribute.String("CommentID", commentRequest.CommentId)),
		oteltrace.WithAttributes(attribute.String("PostID", commentRequest.PostId)),
		oteltrace.WithAttributes(attribute.String("AuthorID", commentRequest.AuthorId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в posts для удаления Comment по ID")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"DeletePostComment_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	err := gc.CommentGRPCService.DeleteComment(
		traceCtx,
		uuid.MustParse(commentRequest.GetCommentId()),
		uuid.MustParse(commentRequest.GetPostId()),
		uuid.MustParse(commentRequest.GetAuthorId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &postsgrpc.DeleteCommentResponse{
		SuccessMessage: "Удаление прошло успешно",
	}, nil
}
