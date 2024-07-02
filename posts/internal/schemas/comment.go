package schemas

import "github.com/google/uuid"

type CommentResponse struct {
	ID       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	AuthorID uuid.UUID `json:"author_id"`
}
