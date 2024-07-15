package likes

import "github.com/google/uuid"

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ReactionResponse struct {
	Id uuid.UUID `json:"id"`
}
