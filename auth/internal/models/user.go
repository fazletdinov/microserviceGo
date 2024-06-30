package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primarykey; type:uuid; default:uuid_generate_v4()"`
	Email     string    `gorm:"not null; unique; size:256; index:user_email_idx, type:btree"`
	FirstName *string   `gorm:"size:256"`
	LastName  *string   `gorm:"size:256"`
	Password  string    `gorm:"not null"`
}
