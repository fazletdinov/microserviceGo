package dto

import "github.com/google/uuid"

type ReactionResponse struct {
	ID       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	PostID   uuid.UUID `json:"post_id"`
}
