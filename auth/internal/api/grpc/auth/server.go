package auth

import (
	authv1 "auth/protogen/auth"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"auth/internal/domain/service"
)

type GRPCController struct {
	authv1.UnimplementedAuthServer
	authService service.GRPCService
	pathSecret  string
}

func Register(gRPC *grpc.Server, authService service.GRPCService, pathSecret string) {
	authv1.RegisterAuthServer(gRPC, &GRPCController{authService: authService, pathSecret: pathSecret})
}

func (s *GRPCController) ExtractUserIDFromToken(
	ctx context.Context,
	authRequest *authv1.ExtractUserIDRequest,
) (*authv1.ExtractUserIDResponse, error) {
	if authRequest.GetAccessToken() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле access_token обязательно")
	}

	userID, err := s.authService.ExtractUserIDFromToken(ctx, authRequest.GetAccessToken(), s.pathSecret)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authv1.ExtractUserIDResponse{
		UserId: userID.String(),
	}, nil
}
