package posts

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

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostResponse struct {
	ID        uuid.UUID          `json:"id"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	AuthorID  uuid.UUID          `json:"author_id"`
	CreatedAt time.Time          `json:"created_at"`
	Comment   *[]CommentResponse `json:"comments"`
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateCommentRequest struct {
	Text string `json:"text"`
}

type CommentResponse struct {
	ID       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	AuthorID uuid.UUID `json:"author_id"`
}

type UpdateCommentRequest struct {
	Text string `json:"text"`
}
