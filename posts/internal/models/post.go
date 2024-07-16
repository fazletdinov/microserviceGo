package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"primarykey; type:uuid;default:uuid_generate_v4()" json:"id"`
	Title     string    `gorm:"not null; size:256" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	AuthorID  uuid.UUID `json:"author_id"`
	Comments  []Comment `gorm:"foreignKey:PostID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
