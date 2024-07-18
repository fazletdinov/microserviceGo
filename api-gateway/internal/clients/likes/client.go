package likes

import (
	likesgrpc "api-grpc-gateway/protogen/golang/likes"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClientLikes struct {
	likes likesgrpc.GatewayLikesClient
}

func NewGRPCClientLikes(
	addrs string,
) (*GRPCClientLikes, error) {
	cc, err := grpc.NewClient(addrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &GRPCClientLikes{likes: likesgrpc.NewGatewayLikesClient(cc)}, nil
}

func (gc *GRPCClientLikes) CreateReaction(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	reactionResponse, err := gc.likes.CreateReaction(ctx, &likesgrpc.CreateReactionRequest{
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
	reactionResponse, err := gc.likes.GetReactionByID(ctx, &likesgrpc.GetReactionRequest{
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
	reactionResponse, err := gc.likes.GetReactions(ctx, &likesgrpc.GetReactionsRequest{
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
	reactionResponse, err := gc.likes.DeleteReaction(ctx, &likesgrpc.DeleteReactionRequest{
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
	reactionResponse, err := gc.likes.DeleteReactionsByAuthor(ctx, &likesgrpc.DeleteReactionsByAuthorRequest{
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
	reactionResponse, err := gc.likes.DeleteReactionsByPost(ctx, &likesgrpc.DeleteReactionsByPostRequest{
		PostId: postID.String(),
	})
	if err != nil {
		return nil, err
	}

	return reactionResponse, nil
}
