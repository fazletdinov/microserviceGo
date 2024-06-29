package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reaction struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primarykey; type:uuid;default:uuid_generate_v4()"`
	AuthorID uuid.UUID `gorm:"not null; unique:author_post"`
	PostID   uuid.UUID `gorm:"not null; unique:author_post"`
}
