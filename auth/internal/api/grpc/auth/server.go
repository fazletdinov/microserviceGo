package auth

import (
	"auth/config"
	authgrpc "auth/protogen/auth"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"auth/internal/domain/service/grpc_service"
)

type GRPCController struct {
	authgrpc.UnimplementedGatewayAuthServer
	UserGRPCService grpcservice.UserGRPCService
	Env             *config.Config
}

func (gc *GRPCController) CreateUser(
	ctx context.Context,
	authRequest *authgrpc.CreateUserRequest,
) (*authgrpc.CreateUserResponse, error) {
	if authRequest.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле email обязательно")
	}
	if authRequest.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле password обязательно")
	}

	userID, err := gc.UserGRPCService.CreateUser(ctx, authRequest.GetEmail(), authRequest.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.CreateUserResponse{
		UserId: userID.String(),
	}, nil
}

func (gc *GRPCController) GetUserByID(
	ctx context.Context,
	authRequest *authgrpc.GetUserRequest,
) (*authgrpc.GetUserResponse, error) {
	if authRequest.GetUserId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле user_id обязательно")
	}

	user, err := gc.UserGRPCService.GetUserByID(ctx, uuid.MustParse(authRequest.GetUserId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.GetUserResponse{
		UserId:    user.ID.String(),
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}, nil
}

func (gc *GRPCController) UpdateUser(
	ctx context.Context,
	authRequest *authgrpc.UpdateUserRequest,
) (*authgrpc.UpdateUserResponse, error) {
	if authRequest.GetUserId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле user_id обязательно")
	}
	if authRequest.GetFirstName() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле first_name обязательно")
	}
	if authRequest.GetLastName() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле last_name обязательно")
	}

	err := gc.UserGRPCService.UpdateUser(ctx, uuid.MustParse(authRequest.GetUserId()), authRequest.GetFirstName(), authRequest.GetLastName())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.UpdateUserResponse{
		SuccessMessage: "Пользователь успешно обновлен",
	}, nil
}

func (gc *GRPCController) DeleteUser(
	ctx context.Context,
	authRequest *authgrpc.DeleteUserRequest,
) (*authgrpc.DeleteUserResponse, error) {
	if authRequest.GetUserId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле user_id обязательно")
	}

	err := gc.UserGRPCService.DeleteUser(ctx, uuid.MustParse(authRequest.GetUserId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.DeleteUserResponse{
		SuccessMessage: "Успешно удалено",
	}, nil
}

func (gc *GRPCController) GetUserByEmail(
	ctx context.Context,
	authRequest *authgrpc.GetUserByEmailRequest,
) (*authgrpc.GetUserResponse, error) {
	if authRequest.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле email обязательно")
	}

	user, err := gc.UserGRPCService.GetUserByEmail(ctx, authRequest.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.GetUserResponse{
		UserId:    user.ID.String(),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}, nil
}

func (gc *GRPCController) GetUserByEmailIsActive(
	ctx context.Context,
	authRequest *authgrpc.GetUserByEmailRequest,
) (*authgrpc.GetUserResponse, error) {
	if authRequest.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле email обязательно")
	}

	user, err := gc.UserGRPCService.GetUserByEmailIsActive(ctx, authRequest.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authgrpc.GetUserResponse{
		UserId:    user.ID.String(),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}, nil
}
