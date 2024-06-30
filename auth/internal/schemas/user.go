package schemas

import (
	"auth/internal/models"
	"context"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName *string   `json:"first_name,omitempty"`
	LastName  *string   `json:"last_name,omitempty"`
	CreatedAt time.Time `json:"create_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserRepository interface {
	Create(c context.Context, user *models.Users) error
	GetByEmail(c context.Context, email string) (models.Users, error)
	GetByID(c context.Context, id uuid.UUID) (models.Users, error)
}
