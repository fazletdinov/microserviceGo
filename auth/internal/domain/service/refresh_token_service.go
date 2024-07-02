package service

import (
	"auth/internal/domain/repository"
	"auth/internal/models"
	"context"

	"github.com/google/uuid"
)

type refreshTokenService struct {
	userRepository repository.UserRepository
}

func NewRefreshTokenService(userRepository repository.UserRepository) RefreshTokenService {
	return &refreshTokenService{
		userRepository: userRepository,
	}
}

func (rtu *refreshTokenService) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.Users, error) {
	return rtu.userRepository.GetByID(ctx, userID)
}

func (rtu *refreshTokenService) CreateAccessToken(user *models.Users, pathSecret string, expiry int) (accessToken string, err error) {
	return GenerateAccessToken(user, pathSecret, expiry)
}

func (rtu *refreshTokenService) CreateRefreshToken(user *models.Users, pathSecret string, expiry int) (refreshToken string, err error) {
	return GenerateRefreshToken(user, pathSecret, expiry)
}

func (rtu *refreshTokenService) ExtractIDFromToken(requestToken string, pathSecret string) (uuid.UUID, error) {
	return ExtractUserIDFromToken(requestToken, pathSecret)
}
