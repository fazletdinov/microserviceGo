package repository

import (
	"auth/internal/models"
	"auth/internal/schemas"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.Users) error
	GetByEmail(ctx context.Context, email string) (*models.Users, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Users, error)
	Update(ctx context.Context, userUpdate *schemas.UpdateUser) (*models.Users, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
