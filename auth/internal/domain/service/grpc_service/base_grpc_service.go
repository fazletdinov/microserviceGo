package grpcservice

import (
	"auth/internal/models"
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, email string, password string) (uuid.UUID, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*models.Users, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, firstName string, lastName string) error
	GetUserByEmail(ctx context.Context, email string) (*models.Users, error)
	GetUserByEmailIsActive(ctx context.Context, email string) (*models.Users, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}
