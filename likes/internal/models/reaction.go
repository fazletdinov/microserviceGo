package models

import (
	"github.com/google/uuid"
	"time"
)

type Reaction struct {
	ID        uuid.UUID `gorm:"primarykey; type:uuid;default:uuid_generate_v4()" json:"id"`
	AuthorID  uuid.UUID `gorm:"not null; unique:author_post" json:"author_id"`
	PostID    uuid.UUID `gorm:"not null; unique:author_post" json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
