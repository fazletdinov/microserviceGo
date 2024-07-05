package service

import (
	"context"

	"github.com/google/uuid"
)

type grpcService struct{}

func NewGRPCService() GRPCService {
	return &grpcService{}
}

func (su *grpcService) ExtractUserIDFromToken(ctx context.Context, accessToken string, pathSecret string) (uuid.UUID, error) {
	return ExtractUserIDFromToken(accessToken, pathSecret)
}
