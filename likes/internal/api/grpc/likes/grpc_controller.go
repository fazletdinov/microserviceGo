package likes

import (
	"context"
	"likes/config"
	likesgrpc "likes/protogen/likes"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	reactionService "likes/internal/domain/service"
)

type GRPCController struct {
	likesgrpc.UnimplementedGatewayLikesServer
	ReactionGRPCService reactionService.ReactionGRPCService
	Env                 *config.Config
}

func (gc *GRPCController) CreateReaction(
	ctx context.Context,
	reactionRequest *likesgrpc.CreateReactionRequest,
) (*likesgrpc.CreateReactionResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if reactionRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if reactionRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"CreateReaction",
		oteltrace.WithAttributes(attribute.String("PostID", reactionRequest.PostId)),
		oteltrace.WithAttributes(attribute.String("AuthorID", reactionRequest.AuthorId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в likes для создания Reaction")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"CreateReaction_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	reactionID, err := gc.ReactionGRPCService.CreateReaction(
		traceCtx,
		uuid.MustParse(reactionRequest.GetPostId()),
		uuid.MustParse(reactionRequest.GetAuthorId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &likesgrpc.CreateReactionResponse{
		ReactionId: reactionID.String(),
	}, nil
}

func (gc *GRPCController) GetReactionByID(
	ctx context.Context,
	reactionRequest *likesgrpc.GetReactionRequest,
) (*likesgrpc.GetReactionResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if reactionRequest.GetReactionId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле reaction_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"GetReactionByID",
		oteltrace.WithAttributes(attribute.String("ReactionID", reactionRequest.ReactionId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в likes для получения Reaction по ID")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"GetReactionByID_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	reaction, err := gc.ReactionGRPCService.GetByID(
		traceCtx,
		uuid.MustParse(reactionRequest.GetReactionId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &likesgrpc.GetReactionResponse{
		ReactionId: reaction.ID.String(),
	}, nil
}

func (gc *GRPCController) GetReactions(
	ctx context.Context,
	reactionRequest *likesgrpc.GetReactionsRequest,
) (*likesgrpc.GetReactionsResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if reactionRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"GetReactions",
		oteltrace.WithAttributes(attribute.String("PostID", reactionRequest.PostId)),
		oteltrace.WithAttributes(attribute.Int64("Limit", int64(reactionRequest.Limit))),
		oteltrace.WithAttributes(attribute.Int64("Offset", int64(reactionRequest.Offset))),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в likes для получения Reactions")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"GetReactions_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	reactions, err := gc.ReactionGRPCService.GetReactionsPost(
		traceCtx,
		uuid.MustParse(reactionRequest.PostId),
		reactionRequest.GetLimit(),
		reactionRequest.GetOffset(),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	arrayReaction := make([]*likesgrpc.Reaction, 0, reactionRequest.GetLimit())
	for _, reaction := range *reactions {
		arrayReaction = append(arrayReaction, &likesgrpc.Reaction{
			ReactionId: reaction.ID.String(),
		})
	}
	return &likesgrpc.GetReactionsResponse{
		Reactions: arrayReaction,
	}, nil
}

func (gc *GRPCController) DeleteReaction(
	ctx context.Context,
	reactionRequest *likesgrpc.DeleteReactionRequest,
) (*likesgrpc.DeleteReactionResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if reactionRequest.GetReactionId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле reaction_id обязательно")
	}
	if reactionRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if reactionRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"DeleteReaction",
		oteltrace.WithAttributes(attribute.String("ReactionID", reactionRequest.ReactionId)),
		oteltrace.WithAttributes(attribute.String("PostID", reactionRequest.PostId)),
		oteltrace.WithAttributes(attribute.String("AuthorID", reactionRequest.AuthorId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в likes для удаления Reaction")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"DeleteReaction_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	err := gc.ReactionGRPCService.DeleteReaction(
		traceCtx,
		uuid.MustParse(reactionRequest.GetReactionId()),
		uuid.MustParse(reactionRequest.GetPostId()),
		uuid.MustParse(reactionRequest.GetAuthorId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &likesgrpc.DeleteReactionResponse{
		SuccessMessage: "Удаление прошло успешно",
	}, nil
}

func (gc *GRPCController) DeleteReactionsByAuthor(
	ctx context.Context,
	reactionRequest *likesgrpc.DeleteReactionsByAuthorRequest,
) (*likesgrpc.DeleteReactionsByAuthorResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if reactionRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"DeleteReactionsByAuthor",
		oteltrace.WithAttributes(attribute.String("AuthorID", reactionRequest.AuthorId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в likes для удаления Reaction автора")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"DeleteReactionsByAuthor_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	err := gc.ReactionGRPCService.DeleteReactionsByAuthor(
		traceCtx,
		uuid.MustParse(reactionRequest.GetAuthorId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &likesgrpc.DeleteReactionsByAuthorResponse{
		SuccessMessage: "Удаление прошло успешно",
	}, nil
}

func (gc *GRPCController) DeleteReactionsByPost(
	ctx context.Context,
	reactionRequest *likesgrpc.DeleteReactionsByPostRequest,
) (*likesgrpc.DeleteReactionsByPostResponse, error) {
	var tracer = otel.Tracer(gc.Env.Jaeger.Application)
	var meter = otel.Meter(gc.Env.Jaeger.Application)

	if reactionRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}

	traceCtx, span := tracer.Start(
		ctx,
		"DeleteReactionsByPost",
		oteltrace.WithAttributes(attribute.String("PostID", reactionRequest.PostId)),
	)
	span.AddEvent("Пришел gRPC запрос от сервиса api-gateway в likes для удаления Reactions Post")
	defer span.End()

	counter, _ := meter.Int64Counter(
		"DeleteReactionsByPost_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	err := gc.ReactionGRPCService.DeleteReactionsByPost(
		traceCtx,
		uuid.MustParse(reactionRequest.GetPostId()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &likesgrpc.DeleteReactionsByPostResponse{
		SuccessMessage: "Удаление прошло успешно",
	}, nil
}
