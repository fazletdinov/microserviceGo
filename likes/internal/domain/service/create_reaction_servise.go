package service

import (
	"context"
	"likes/internal/domain/repository"
	"likes/internal/models"
)

type reactionCreateService struct {
	reactionRepository repository.ReactionRepository
}

func NewCreateReactionService(reactionRepository repository.ReactionRepository) CreateReactionServcie {
	return &reactionCreateService{
		reactionRepository: reactionRepository,
	}
}

func (rs *reactionCreateService) CreateReaction(ctx context.Context, reaction *models.Reaction) error {
	return rs.reactionRepository.Create(ctx, reaction)
}
