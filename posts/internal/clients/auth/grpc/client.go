package grpc

import (
	"context"
	authv1 "posts/protogen/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	api authv1.AuthClient
}

func NewGRPCClient(
	addrs string,
) (*GRPCClient, error) {
	cc, err := grpc.NewClient(addrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &GRPCClient{api: authv1.NewAuthClient(cc)}, nil
}

func (gc *GRPCClient) ExtractUserIDFromToken(ctx context.Context, accessToken string) (string, error) {
	userID, err := gc.api.ExtractUserIDFromToken(ctx, &authv1.ExtractUserIDRequest{
		AccessToken: accessToken,
	})

	if err != nil {
		return "", err
	}

	return userID.UserId, nil
}
