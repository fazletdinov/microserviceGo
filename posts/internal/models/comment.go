package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primarykey; type:uuid;default:uuid_generate_v4()"`
	Text     string    `gorm:"not null; size:256"`
	AuthorID uuid.UUID
	PostID   uuid.UUID
}
