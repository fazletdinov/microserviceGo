package grpcservice

import (
	"auth/internal/dto"
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, email string, password string) (uuid.UUID, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, firstName string, lastName string) error
	GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error)
	GetUserByEmailIsActive(ctx context.Context, email string) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}
