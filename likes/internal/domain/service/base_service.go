package service

import (
	"context"
	"likes/internal/dto"

	"github.com/google/uuid"
)

type ReactionService interface {
	CreateReaction(ctx context.Context, postID uuid.UUID, authorID uuid.UUID) (uuid.UUID, error)
	GetByID(ctx context.Context, reactionID uuid.UUID) (*dto.ReactionResponse, error)
	GetReactionsPost(ctx context.Context, postID uuid.UUID, limit uint64, offset uint64) (*[]dto.ReactionResponse, error)
	DeleteReaction(ctx context.Context, reactionID uuid.UUID, postID uuid.UUID, authorID uuid.UUID) error
	DeleteReactionsByAuthor(ctx context.Context, authorID uuid.UUID) error
	DeleteReactionsByPost(ctx context.Context, postID uuid.UUID) error
}
