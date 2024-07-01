package service

import (
	"auth/internal/domain/repository"
	"auth/internal/models"

	"context"

	"github.com/google/uuid"
)

type deleteService struct {
	userRepository repository.UserRepository
}

func NewDeleteService(userRepository repository.UserRepository) DeleteService {
	return &deleteService{
		userRepository: userRepository,
	}
}

func (ds *deleteService) GetUserByID(ctx context.Context, id uuid.UUID) (*models.Users, error) {
	return ds.userRepository.GetByID(ctx, id)
}

func (ds *deleteService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return ds.userRepository.Delete(ctx, id)
}
