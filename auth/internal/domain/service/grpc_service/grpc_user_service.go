package grpcservice

import (
	"auth/internal/domain/repository"
	"auth/internal/models"
	"context"

	"github.com/google/uuid"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (ugs *userService) CreateUser(ctx context.Context, email string, password string) (uuid.UUID, error) {
	return ugs.userRepository.Create(ctx, email, password)
}

func (ugs *userService) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.Users, error) {
	return ugs.userRepository.GetByID(ctx, userID)
}

func (ugs *userService) UpdateUser(ctx context.Context, userID uuid.UUID, firstName string, lastName string) error {
	return ugs.userRepository.Update(ctx, userID, firstName, lastName)
}

func (ugs *userService) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	return ugs.userRepository.GetByEmail(ctx, email)
}

func (ugs *userService) GetUserByEmailIsActive(ctx context.Context, email string) (*models.Users, error) {
	return ugs.userRepository.GetByEmailIsActive(ctx, email)
}

func (ugs *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return ugs.userRepository.Delete(ctx, userID)
}
