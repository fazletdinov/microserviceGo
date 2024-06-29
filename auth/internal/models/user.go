package models

import (
	"auth/internal/database/postgres"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primarykey; type:uuid; default:uuid_generate_v4()"`
	Email     string    `gorm:"not null; unique; size:256; index:user_email_idx, type:btree"`
	FirstName *string   `gorm:"size:256"`
	LastName  *string   `gorm:"size:256"`
	Password  string    `gorm:"not null; hash" json:"-"`
}

func (user *User) Save() (*User, error) {
	err := postgres.DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	return nil
}
