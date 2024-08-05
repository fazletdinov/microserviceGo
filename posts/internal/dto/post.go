package dto

import (
	"posts/internal/models"
	"time"

	"github.com/google/uuid"
)

type PostResponse struct {
	ID        uuid.UUID         `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	AuthorID  uuid.UUID         `json:"author_id"`
	CreatedAt time.Time         `json:"created_at"`
	Comments  *[]models.Comment `json:"comments"`
}
