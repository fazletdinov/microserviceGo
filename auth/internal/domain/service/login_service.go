package service

import (
	"auth/internal/domain/repository"
	"auth/internal/models"
	"context"
)

type loginService struct {
	userRepository repository.UserRepository
}

func NewLoginService(userRepository repository.UserRepository) LoginService {
	return &loginService{
		userRepository: userRepository,
	}
}

func (lc *loginService) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	return lc.userRepository.GetByEmail(ctx, email)
}

func (lu *loginService) CreateAccessToken(user *models.Users, pathSecret string, expiry int) (accessToken string, err error) {
	return GenerateAccessToken(user, pathSecret, expiry)
}

func (lu *loginService) CreateRefreshToken(user *models.Users, pathSecret string, expiry int) (refreshToken string, err error) {
	return GenerateRefreshToken(user, pathSecret, expiry)
}
