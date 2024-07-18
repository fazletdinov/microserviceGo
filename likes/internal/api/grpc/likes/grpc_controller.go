package likes

import (
	"context"
	"likes/config"
	likesgrpc "likes/protogen/likes"

	"github.com/google/uuid"
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
	if reactionRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if reactionRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}

	reactionID, err := gc.ReactionGRPCService.CreateReaction(
		ctx,
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
	if reactionRequest.GetReactionId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле reaction_id обязательно")
	}

	reaction, err := gc.ReactionGRPCService.GetByID(
		ctx,
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
	reactions, err := gc.ReactionGRPCService.GetReactionsPost(
		ctx,
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
	if reactionRequest.GetReactionId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле reaction_id обязательно")
	}
	if reactionRequest.GetPostId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле post_id обязательно")
	}
	if reactionRequest.GetAuthorId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле author_id обязательно")
	}
	err := gc.ReactionGRPCService.DeleteReaction(
		ctx,
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
