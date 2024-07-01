package schemas

import (
	"github.com/google/uuid"
)

type UpdateUser struct {
	ID        uuid.UUID `json:"-"`
	FirstName *string   `json:"first_name"`
	LastName  *string   `json:"last_name"`
}
