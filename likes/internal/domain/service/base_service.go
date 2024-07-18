package service

import (
	"context"
	"likes/internal/models"

	"github.com/google/uuid"
)

type ReactionGRPCService interface {
	CreateReaction(ctx context.Context, postID uuid.UUID, authorID uuid.UUID) (uuid.UUID, error)
	GetByID(ctx context.Context, reactionID uuid.UUID) (*models.Reaction, error)
	GetReactionsPost(ctx context.Context, postID uuid.UUID, limit uint64, offset uint64) (*[]models.Reaction, error)
	DeleteReaction(ctx context.Context, reactionID uuid.UUID, postID uuid.UUID, authorID uuid.UUID) error
}
