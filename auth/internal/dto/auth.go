package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName *string   `json:"first_name,omitempty"`
	LastName  *string   `json:"last_name,omitempty"`
	CreatedAt time.Time `json:"create_at,omitempty"`
}
