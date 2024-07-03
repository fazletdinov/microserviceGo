package schemas

import (
	"github.com/google/uuid"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ReactionCreateRequest struct {
	PostID   uuid.UUID `json:"-"`
	AuthorID uuid.UUID `json:"author_id"`
}

type ReactionResponse struct {
	ID       uuid.UUID `json:"id"`
	PostID   uuid.UUID `json:"post_id"`
	AuthorID uuid.UUID `json:"author_id"`
}
