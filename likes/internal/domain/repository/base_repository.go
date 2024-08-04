package repository

import (
	"context"
	"likes/internal/models"

	"github.com/google/uuid"
)

type ReactionRepository interface {
	Create(ctx context.Context, postID uuid.UUID, authorID uuid.UUID) (uuid.UUID, error)
	GetByIDReaction(ctx context.Context, reactionID uuid.UUID) (*models.Reaction, error)
	GetReactionsPost(ctx context.Context, postID uuid.UUID, limit uint64, offset uint64) (*[]models.Reaction, error)
	DeleteReaction(ctx context.Context, reactionID uuid.UUID, postID uuid.UUID, authorID uuid.UUID) error
	DeleteReactionsByAuthor(ctx context.Context, authorID uuid.UUID) error
	DeleteReactionsByPost(ctx context.Context, postID uuid.UUID) error
}
