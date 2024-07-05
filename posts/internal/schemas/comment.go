package schemas

import "github.com/google/uuid"

type CommentResponse struct {
	ID       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	AuthorID uuid.UUID `json:"author_id"`
}

type CommentCreateRequest struct {
	Text     string    `json:"text"`
	AuthorID uuid.UUID `json:"-"`
	PostID   uuid.UUID `json:"-"`
}

type CommentUpdateRequest struct {
	Text string `json:"text"`
}
