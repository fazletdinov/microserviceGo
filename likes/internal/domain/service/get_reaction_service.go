package service

import (
	"context"
	"likes/internal/domain/repository"
	"likes/internal/models"

	"github.com/google/uuid"
)

type getReactionService struct {
	reactionRepository repository.ReactionRepository
}

func NewGetReactionService(reactionRepository repository.ReactionRepository) GetReactionServcie {
	return &getReactionService{
		reactionRepository: reactionRepository,
	}
}

func (rc *getReactionService) GetByID(ctx context.Context, postID uuid.UUID) (*models.Reaction, error) {
	return rc.reactionRepository.GetByIDReaction(ctx, postID)
}

func (rc *getReactionService) GetReactionsPost(ctx context.Context, postID uuid.UUID, limit int, offset int) (*[]models.Reaction, error) {
	return rc.reactionRepository.GetReactionsPost(ctx, postID, limit, offset)
}
