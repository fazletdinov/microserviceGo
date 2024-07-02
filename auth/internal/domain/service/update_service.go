package service

import (
	"auth/internal/domain/repository"
	"auth/internal/schemas"
	"context"

	"auth/internal/models"

	"github.com/google/uuid"
)

type updateService struct {
	userRepository repository.UserRepository
}

func NewUpdateService(userRepository repository.UserRepository) UpdateService {
	return &updateService{
		userRepository: userRepository,
	}
}

func (up *updateService) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.Users, error) {
	return up.userRepository.GetByID(ctx, userID)
}

func (up *updateService) UpdateUser(ctx context.Context, userID uuid.UUID, userUpdate *schemas.UpdateUser) error {
	err := up.userRepository.Update(ctx, userID, userUpdate)
	if err != nil {
		return err
	}
	return nil
}
