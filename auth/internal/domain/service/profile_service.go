package service

import (
	"auth/internal/schemas"
	"context"

	"github.com/google/uuid"
)

type profileService struct {
	userRepository schemas.UserRepository
}

func NewProfileService(userRepository schemas.UserRepository) schemas.ProfileService {
	return &profileService{
		userRepository: userRepository,
	}
}

func (ps *profileService) GetProfileByID(ctx context.Context, userID uuid.UUID) (*schemas.UserResponse, error) {
	user, err := ps.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	profile := schemas.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
	}

	return &profile, nil
}
