package service

import (
	"auth/internal/models"
	"auth/internal/schemas"
	"context"

	"github.com/google/uuid"
)

type LoginService interface {
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	CreateAccessToken(user *models.Users, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *models.Users, secret string, expiry int) (refreshToken string, err error)
}

type ProfileService interface {
	GetProfileByID(ctx context.Context, userID uuid.UUID) (*schemas.UserResponse, error)
}

type RefreshTokenService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.Users, error)
	CreateAccessToken(user *models.Users, pathSecret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *models.Users, pathSecret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, pathSecret string) (uuid.UUID, error)
}

type SignupService interface {
	Create(ctx context.Context, user *models.Users) error
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	CreateAccessToken(user *models.Users, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *models.Users, secret string, expiry int) (refreshToken string, err error)
}

type UpdateService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.Users, error)
	UpdateUser(ctx context.Context, updateUser *schemas.UpdateUser) (*schemas.UserResponse, error)
}

type DeleteService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.Users, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
