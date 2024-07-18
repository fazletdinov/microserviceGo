package service

import (
	"context"
	"likes/internal/domain/repository"
	"likes/internal/models"

	"github.com/google/uuid"
)

type reactionGRPCService struct {
	reactionGRPCRepository repository.ReactionGRPCRepository
}

func NewReactionGRPCService(
	reactionGRPCRepository repository.ReactionGRPCRepository,
) ReactionGRPCService {
	return &reactionGRPCService{
		reactionGRPCRepository: reactionGRPCRepository,
	}
}

func (rs *reactionGRPCService) CreateReaction(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	return rs.reactionGRPCRepository.Create(ctx, postID, authorID)
}

func (rs *reactionGRPCService) GetByID(
	ctx context.Context,
	reactionID uuid.UUID,
) (*models.Reaction, error) {
	return rs.reactionGRPCRepository.GetByIDReaction(ctx, reactionID)
}

func (rs *reactionGRPCService) GetReactionsPost(
	ctx context.Context,
	postID uuid.UUID,
	limit uint64,
	offset uint64,
) (*[]models.Reaction, error) {
	return rs.reactionGRPCRepository.GetReactionsPost(ctx, postID, limit, offset)
}

func (rs *reactionGRPCService) DeleteReaction(
	ctx context.Context,
	reactionID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	return rs.reactionGRPCRepository.DeleteReaction(ctx, reactionID, postID, authorID)
}
