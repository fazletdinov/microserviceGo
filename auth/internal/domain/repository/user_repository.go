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

func NewUserRepository(db *gorm.DB) UserRepository {
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

func (ur *userRepository) GetByEmail(c context.Context, email string) (*models.Users, error) {
	var user models.Users
	result := ur.database.Where("email = ? AND is_active = ?", email, true).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *userRepository) GetByID(c context.Context, userID uuid.UUID) (*models.Users, error) {
	var user models.Users
	result := ur.database.Where("id = ? AND is_active = ?", userID, true).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *userRepository) Update(c context.Context, userID uuid.UUID, userUpdate *schemas.UpdateUser) error {
	result := ur.database.Model(&models.Users{}).Where("is_active = ? AND id = ?", true, userID).Updates(&userUpdate)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepository) Delete(c context.Context, userID uuid.UUID) error {
	result := ur.database.Model(&models.Users{}).Where("id = ?", userID).Update("is_active", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
