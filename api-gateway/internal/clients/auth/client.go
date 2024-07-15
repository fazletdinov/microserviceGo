package auth

import (
	authgrpc "api-grpc-gateway/protogen/golang/auth"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClientAuth struct {
	auth authgrpc.GatewayAuthClient
}

func NewGRPCClientAuth(
	addrs string,
) (*GRPCClientAuth, error) {
	cc, err := grpc.NewClient(addrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &GRPCClientAuth{auth: authgrpc.NewGatewayAuthClient(cc)}, nil
}

func (gc *GRPCClientAuth) CreateUser(
	ctx context.Context,
	email string,
	password string,
) (uuid.UUID, error) {
	response, err := gc.auth.CreateUser(ctx, &authgrpc.CreateUserRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(response.UserId), nil
}

func (gc *GRPCClientAuth) GetUserByID(
	ctx context.Context,
	userID uuid.UUID,
) (*authgrpc.GetUserResponse, error) {
	response, err := gc.auth.GetUserByID(ctx, &authgrpc.GetUserRequest{
		UserId: userID.String(),
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (gc *GRPCClientAuth) UpdateUser(
	ctx context.Context,
	userID uuid.UUID,
	firstName string,
	lastName string,
) (string, error) {
	response, err := gc.auth.UpdateUser(ctx, &authgrpc.UpdateUserRequest{
		UserId:    userID.String(),
		FirstName: firstName,
		LastName:  lastName,
	})

	if err != nil {
		return "", err
	}

	return response.SuccessMessage, nil
}

func (gc *GRPCClientAuth) DeleteUser(ctx context.Context, userID uuid.UUID) (string, error) {
	response, err := gc.auth.DeleteUser(ctx, &authgrpc.DeleteUserRequest{
		UserId: userID.String(),
	})

	if err != nil {
		return "", err
	}

	return response.SuccessMessage, nil
}

func (gc *GRPCClientAuth) GetUserByEmail(ctx context.Context, email string) (*authgrpc.GetUserResponse, error) {
	response, err := gc.auth.GetUserByEmail(ctx, &authgrpc.GetUserByEmailRequest{
		Email: email,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (gc *GRPCClientAuth) GetUserByEmailIsActive(ctx context.Context, email string) (*authgrpc.GetUserResponse, error) {
	response, err := gc.auth.GetUserByEmailIsActive(ctx, &authgrpc.GetUserByEmailRequest{
		Email: email,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
