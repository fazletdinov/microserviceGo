package schemas

import (
	"time"

	"github.com/google/uuid"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type PostCreateRequest struct {
	ID       uuid.UUID `json:"-"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	AuthorID uuid.UUID `json:"-"`
}

type PostUpdatedResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  uuid.UUID `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PostResponse struct {
	ID        uuid.UUID          `json:"id"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	AuthorID  uuid.UUID          `json:"author_id"`
	CreatedAt time.Time          `json:"created_at"`
	Comment   *[]CommentResponse `json:"comments"`
}

type PostUpdateRequest struct {
	ID       uuid.UUID `json:"-"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	AuthorID uuid.UUID `json:"-"`
}
