package service

import (
	"auth/internal/models"
	"auth/internal/schemas"
	"context"
)

type signupService struct {
	userRepository schemas.UserRepository
}

func NewSignupService(userRepository schemas.UserRepository) schemas.SignupService {
	return &signupService{
		userRepository: userRepository,
	}
}

func (su *signupService) Create(ctx context.Context, user *models.Users) error {
	return su.userRepository.Create(ctx, user)
}

func (su *signupService) GetUserByEmail(ctx context.Context, email string) (models.Users, error) {
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupService) CreateAccessToken(user *models.Users, pathSecret string, expiry int) (accessToken string, err error) {
	return GenerateAccessToken(user, pathSecret, expiry)
}

func (su *signupService) CreateRefreshToken(user *models.Users, pathSecret string, expiry int) (refreshToken string, err error) {
	return GenerateRefreshToken(user, pathSecret, expiry)
}
