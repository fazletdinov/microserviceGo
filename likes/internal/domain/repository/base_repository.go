package repository

import (
	"context"
	"likes/internal/models"

	"github.com/google/uuid"
)

type ReactionRepository interface {
	Create(ctx context.Context, reaction *models.Reaction) error
	GetByIDReaction(ctx context.Context, reactionID uuid.UUID) (*models.Reaction, error)
	GetReactionsPost(ctx context.Context, postID uuid.UUID, limit int, offset int) (*[]models.Reaction, error)
	DeleteReaction(ctx context.Context, reactionID uuid.UUID) error
}
