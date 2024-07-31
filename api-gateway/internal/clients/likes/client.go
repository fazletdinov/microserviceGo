package likes

import (
	likesgrpc "api-grpc-gateway/protogen/golang/likes"
	"context"

	"api-grpc-gateway/config"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClientLikes struct {
	likes likesgrpc.GatewayLikesClient
	env   *config.Config
}

func NewGRPCClientLikes(
	addrs string,
	env *config.Config,
) (*GRPCClientLikes, error) {
	cc, err := grpc.NewClient(addrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &GRPCClientLikes{
			likes: likesgrpc.NewGatewayLikesClient(cc),
			env:   env,
		},
		nil
}

func (gc *GRPCClientLikes) CreateReaction(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"CreateReaction",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис likes для создания реакции на post")
	defer span.End()
	reactionResponse, err := gc.likes.CreateReaction(traceCtx, &likesgrpc.CreateReactionRequest{
		PostId:   postID.String(),
		AuthorId: authorID.String(),
	})

	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(reactionResponse.ReactionId), nil
}

func (gc *GRPCClientLikes) GetReactionByID(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*likesgrpc.GetReactionResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"GetReactionByID",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис likes для получения реакции на post")
	defer span.End()
	reactionResponse, err := gc.likes.GetReactionByID(traceCtx, &likesgrpc.GetReactionRequest{
		PostId:   postID.String(),
		AuthorId: authorID.String(),
	})
	if err != nil {
		return nil, err
	}

	return reactionResponse, nil
}

func (gc *GRPCClientLikes) GetReactionsPost(
	ctx context.Context,
	postID uuid.UUID,
	limit uint64,
	offset uint64,
) (*likesgrpc.GetReactionsResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"GetReactionsPost",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.Int64("Limit", int64(limit))),
		oteltrace.WithAttributes(attribute.Int64("Offset", int64(offset))),
	)
	span.AddEvent("Начало gRPC запроса в сервис likes для получения списка реакции на post")
	defer span.End()
	reactionResponse, err := gc.likes.GetReactions(traceCtx, &likesgrpc.GetReactionsRequest{
		PostId: postID.String(),
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	return reactionResponse, nil
}

func (gc *GRPCClientLikes) DeleteReaction(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*likesgrpc.DeleteReactionResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"DeleteReaction",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис likes для удаления реакции")
	defer span.End()
	reactionResponse, err := gc.likes.DeleteReaction(traceCtx, &likesgrpc.DeleteReactionRequest{
		AuthorId: authorID.String(),
		PostId:   postID.String(),
	})
	if err != nil {
		return nil, err
	}

	return reactionResponse, nil
}

func (gc *GRPCClientLikes) DeleteReactionsByAuthor(
	ctx context.Context,
	authorID uuid.UUID,
) (*likesgrpc.DeleteReactionsByAuthorResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"DeleteReactionsByAuthor",
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис likes для удаления реакций на post по AuthorID")
	defer span.End()
	reactionResponse, err := gc.likes.DeleteReactionsByAuthor(traceCtx, &likesgrpc.DeleteReactionsByAuthorRequest{
		AuthorId: authorID.String(),
	})
	if err != nil {
		return nil, err
	}

	return reactionResponse, nil
}

func (gc *GRPCClientLikes) DeleteReactionsByPost(
	ctx context.Context,
	postID uuid.UUID,
) (*likesgrpc.DeleteReactionsByPostResponse, error) {
	var tracer = otel.Tracer(gc.env.Jaeger.ServerName)
	traceCtx, span := tracer.Start(
		ctx,
		"DeleteReactionsByPost",
		oteltrace.WithAttributes(attribute.String("PostID", postID.String())),
	)
	span.AddEvent("Начало gRPC запроса в сервис likes для удаления реакций на post")
	defer span.End()
	reactionResponse, err := gc.likes.DeleteReactionsByPost(traceCtx, &likesgrpc.DeleteReactionsByPostRequest{
		PostId: postID.String(),
	})
	if err != nil {
		return nil, err
	}

	return reactionResponse, nil
}
