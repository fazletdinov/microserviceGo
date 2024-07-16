package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `gorm:"primarykey; type:uuid;default:uuid_generate_v4()" json:"id"`
	Text      string    `gorm:"not null; size:256" json:"text"`
	AuthorID  uuid.UUID `json:"author_id"`
	PostID    uuid.UUID `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
