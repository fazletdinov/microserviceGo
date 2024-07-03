package service

import (
	"context"
	"likes/internal/models"

	"github.com/google/uuid"
)

type CreateReactionServcie interface {
	CreateReaction(ctx context.Context, reaction *models.Reaction) error
}

type DeleteReactionServcie interface {
	GetByID(ctx context.Context, reactionID uuid.UUID) (*models.Reaction, error)
	DeleteReaction(ctx context.Context, reactionID uuid.UUID) error
}

type GetReactionServcie interface {
	GetByID(ctx context.Context, reactionID uuid.UUID) (*models.Reaction, error)
	GetReactionsPost(ctx context.Context, postID uuid.UUID, limit int, offset int) (*[]models.Reaction, error)
}
