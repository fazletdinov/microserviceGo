package grpcservice

import (
	"auth/internal/domain/repository"
	"auth/internal/models"
	"context"

	"github.com/google/uuid"
)

type userGRPCService struct {
	userGRPCRepository repository.UserGRPCRepository
}

func NewUserGRPCService(userGRPCRepository repository.UserGRPCRepository) UserGRPCService {
	return &userGRPCService{
		userGRPCRepository: userGRPCRepository,
	}
}

func (ugs *userGRPCService) CreateUser(ctx context.Context, email string, password string) (uuid.UUID, error) {
	return ugs.userGRPCRepository.Create(ctx, email, password)
}

func (ugs *userGRPCService) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.Users, error) {
	return ugs.userGRPCRepository.GetByID(ctx, userID)
}

func (ugs *userGRPCService) UpdateUser(ctx context.Context, userID uuid.UUID, firstName string, lastName string) error {
	return ugs.userGRPCRepository.Update(ctx, userID, firstName, lastName)
}

func (ugs *userGRPCService) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	return ugs.userGRPCRepository.GetByEmail(ctx, email)
}

func (ugs *userGRPCService) GetUserByEmailIsActive(ctx context.Context, email string) (*models.Users, error) {
	return ugs.userGRPCRepository.GetByEmailIsActive(ctx, email)
}

func (ugs *userGRPCService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return ugs.userGRPCRepository.Delete(ctx, userID)
}
