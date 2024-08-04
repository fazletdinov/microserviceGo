package repository

import (
	"auth/internal/models"
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

func (ur *userRepository) Create(
	ctx context.Context,
	email string,
	password string,
) (uuid.UUID, error) {
	user := models.Users{
		Email:    email,
		Password: password,
		IsActive: true,
	}
	result := ur.database.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return user.ID, nil
}

func (ur *userRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*models.Users, error) {
	var user models.Users
	result := ur.database.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *userRepository) GetByEmailIsActive(
	ctx context.Context,
	email string,
) (*models.Users, error) {
	var user models.Users
	result := ur.database.WithContext(ctx).Where("email = ? AND is_active = ?", email, true).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *userRepository) GetByID(
	ctx context.Context,
	userID uuid.UUID,
) (*models.Users, error) {
	var user models.Users
	result := ur.database.WithContext(ctx).Where("is_active = ?", true).First(&user, "id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *userRepository) Update(
	ctx context.Context,
	userID uuid.UUID,
	firstName string,
	lastName string,
) error {
	result := ur.database.WithContext(ctx).Model(&models.Users{ID: userID, IsActive: true}).Updates(models.Users{FirstName: &firstName, LastName: &lastName})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepository) Delete(
	ctx context.Context,
	userID uuid.UUID,
) error {
	result := ur.database.WithContext(ctx).Model(&models.Users{}).Where("id = ?", userID).Update("is_active", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
