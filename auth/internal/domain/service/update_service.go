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

func (up *updateService) GetUserByID(ctx context.Context, id uuid.UUID) (*models.Users, error) {
	return up.userRepository.GetByID(ctx, id)
}

func (up *updateService) UpdateUser(ctx context.Context, userUpdate *schemas.UpdateUser) (*schemas.UserResponse, error) {
	user, err := up.userRepository.Update(ctx, userUpdate)
	if err != nil {
		return nil, err
	}
	updatedUser := schemas.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &updatedUser, nil
}
