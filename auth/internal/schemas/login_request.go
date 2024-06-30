package schemas

import (
	"auth/internal/models"
	"context"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginService interface {
	GetUserByEmail(c context.Context, email string) (models.Users, error)
	CreateAccessToken(user *models.Users, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *models.Users, secret string, expiry int) (refreshToken string, err error)
}
