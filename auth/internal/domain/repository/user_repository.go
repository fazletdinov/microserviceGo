package repository

import (
	"auth/internal/models"
	"auth/internal/schemas"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	result := ur.database.Where("email = ? AND is_active = ?", email, true).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *userRepository) GetByID(c context.Context, id uuid.UUID) (*models.Users, error) {
	var user models.Users
	result := ur.database.Where("id = ? AND is_active = ?", id, true).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *userRepository) Update(c context.Context, userUpdate *schemas.UpdateUser) (*models.Users, error) {
	var user models.Users
	result := ur.database.Model(&user).Clauses(clause.Returning{}).Where("is_active = ?", true).Updates(&userUpdate)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *userRepository) Delete(c context.Context, id uuid.UUID) error {
	var user models.Users
	result := ur.database.Model(&user).Where("id = ?", id).Update("is_active", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
