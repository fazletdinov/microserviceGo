package grpcservice

import (
	"auth/internal/domain/repository"
	"auth/internal/dto"
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

func (ugs *userService) GetUserByID(ctx context.Context, userID uuid.UUID) (*dto.UserResponse, error) {
	userResponse, err := ugs.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:        userResponse.ID,
		Email:     userResponse.Email,
		Password:  userResponse.Password,
		FirstName: userResponse.FirstName,
		LastName:  userResponse.LastName,
		CreatedAt: userResponse.CreatedAt,
	}, nil
}
func (ugs *userService) UpdateUser(ctx context.Context, userID uuid.UUID, firstName string, lastName string) error {
	return ugs.userRepository.Update(ctx, userID, firstName, lastName)
}

func (ugs *userService) GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	userResponse, err := ugs.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:        userResponse.ID,
		Email:     userResponse.Email,
		Password:  userResponse.Password,
		FirstName: userResponse.FirstName,
		LastName:  userResponse.LastName,
		CreatedAt: userResponse.CreatedAt,
	}, nil
}

func (ugs *userService) GetUserByEmailIsActive(ctx context.Context, email string) (*dto.UserResponse, error) {
	userResponse, err := ugs.userRepository.GetByEmailIsActive(ctx, email)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:        userResponse.ID,
		Email:     userResponse.Email,
		Password:  userResponse.Password,
		FirstName: userResponse.FirstName,
		LastName:  userResponse.LastName,
		CreatedAt: userResponse.CreatedAt,
	}, nil
}

func (ugs *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return ugs.userRepository.Delete(ctx, userID)
}
