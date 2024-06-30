package service

import (
	"auth/internal/models"
	"auth/internal/schemas"
	"context"

	"github.com/google/uuid"
)

type refreshTokenUsecase struct {
	userRepository schemas.UserRepository
}

func NewRefreshTokenService(userRepository schemas.UserRepository) schemas.RefreshTokenService {
	return &refreshTokenUsecase{
		userRepository: userRepository,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(ctx context.Context, id uuid.UUID) (models.Users, error) {
	return rtu.userRepository.GetByID(ctx, id)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *models.Users, pathSecret string, expiry int) (accessToken string, err error) {
	return GenerateAccessToken(user, pathSecret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *models.Users, pathSecret string, expiry int) (refreshToken string, err error) {
	return GenerateRefreshToken(user, pathSecret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, pathSecret string) (uuid.UUID, error) {
	return ExtractUserIDFromToken(requestToken, pathSecret)
}
