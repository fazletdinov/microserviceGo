package schemas

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	CreatedAt time.Time `json:"create_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ProfileService interface {
	GetProfileByID(c context.Context, userID uuid.UUID) (*UserResponse, error)
}
