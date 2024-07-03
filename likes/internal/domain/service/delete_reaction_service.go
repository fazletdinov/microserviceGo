package service

import (
	"context"
	"likes/internal/domain/repository"
	"likes/internal/models"

	"github.com/google/uuid"
)

type deleteReactionService struct {
	reactionRepository repository.ReactionRepository
}

func NewDeleteReactionService(reactionRepository repository.ReactionRepository) DeleteReactionServcie {
	return &deleteReactionService{
		reactionRepository: reactionRepository,
	}
}

func (drs *deleteReactionService) GetByID(ctx context.Context, reactionID uuid.UUID) (*models.Reaction, error) {
	return drs.reactionRepository.GetByIDReaction(ctx, reactionID)
}

func (drs *deleteReactionService) DeleteReaction(ctx context.Context, reactionID uuid.UUID) error {
	return drs.reactionRepository.DeleteReaction(ctx, reactionID)
}
