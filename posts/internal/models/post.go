package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primarykey; type:uuid;default:uuid_generate_v4()"`
	Title    string    `gorm:"not null; size:256"`
	Content  string    `gorm:"not null"`
	AuthorID uuid.UUID
	Comment  []Comment `gorm:"foreignKey:PostID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
