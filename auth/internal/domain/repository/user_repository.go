package repository

import (
	"auth/internal/models"
	"auth/internal/schemas"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) schemas.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *models.Users) error {
	result := ur.database.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (models.Users, error) {
	var user models.Users
	result := ur.database.First(&user, "email = ?", email)
	if result.Error != nil {
		return models.Users{}, result.Error
	}
	return user, nil
}

func (ur *userRepository) GetByID(c context.Context, id uuid.UUID) (models.Users, error) {
	var user models.Users
	result := ur.database.First(&user, "id = ?", id)
	if result.Error != nil {
		return models.Users{}, result.Error
	}
	return user, nil
}
