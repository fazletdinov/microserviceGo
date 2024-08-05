package dto

import (
	"github.com/google/uuid"
	"time"
)

type CommentResponse struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	AuthorID  uuid.UUID `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
