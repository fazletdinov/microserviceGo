package service

import (
	"auth/internal/models"
	"auth/internal/schemas"
	"context"
)

type loginService struct {
	userRepository schemas.UserRepository
}

func NewLoginService(userRepository schemas.UserRepository) schemas.LoginService {
	return &loginService{
		userRepository: userRepository,
	}
}

func (lc *loginService) GetUserByEmail(ctx context.Context, email string) (models.Users, error) {
	return lc.userRepository.GetByEmail(ctx, email)
}

func (lu *loginService) CreateAccessToken(user *models.Users, pathSecret string, expiry int) (accessToken string, err error) {
	return GenerateAccessToken(user, pathSecret, expiry)
}

func (lu *loginService) CreateRefreshToken(user *models.Users, pathSecret string, expiry int) (refreshToken string, err error) {
	return GenerateRefreshToken(user, pathSecret, expiry)
}
