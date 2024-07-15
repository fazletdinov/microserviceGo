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
	GetByID(ctx context.Context, userID uuid.UUID) (*models.Users, error)
	Update(ctx context.Context, userID uuid.UUID, userUpdate *schemas.UpdateUser) error
	Delete(ctx context.Context, userID uuid.UUID) error
}

type UserGRPCRepository interface {
	Create(ctx context.Context, email string, password string) (uuid.UUID, error)
	GetByEmail(ctx context.Context, email string) (*models.Users, error)
	GetByEmailIsActive(ctx context.Context, email string) (*models.Users, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*models.Users, error)
	Update(ctx context.Context, userID uuid.UUID, firstName string, lastName string) error
	Delete(ctx context.Context, userID uuid.UUID) error
}
