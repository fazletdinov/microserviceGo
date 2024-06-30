package schemas

import (
	"auth/internal/models"
	"context"

	"github.com/google/uuid"
)

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenService interface {
	GetUserByID(c context.Context, id uuid.UUID) (models.Users, error)
	CreateAccessToken(user *models.Users, pathSecret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *models.Users, pathSecret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, pathSecret string) (uuid.UUID, error)
}
