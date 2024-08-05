package service

import (
	"context"
	"likes/internal/domain/repository"
	"likes/internal/dto"

	"github.com/google/uuid"
)

type reactionService struct {
	reactionRepository repository.ReactionRepository
}

func NewReactionService(
	reactionRepository repository.ReactionRepository,
) ReactionService {
	return &reactionService{
		reactionRepository: reactionRepository,
	}
}

func (rs *reactionService) CreateReaction(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	return rs.reactionRepository.Create(ctx, postID, authorID)
}

func (rs *reactionService) GetByID(
	ctx context.Context,
	reactionID uuid.UUID,
) (*dto.ReactionResponse, error) {
	reactionResponse, err := rs.reactionRepository.GetByIDReaction(ctx, reactionID)
	if err != nil {
		return nil, err
	}
	return &dto.ReactionResponse{
		ID:       reactionResponse.ID,
		AuthorID: reactionResponse.AuthorID,
		PostID:   reactionResponse.PostID,
	}, nil
}

func (rs *reactionService) GetReactionsPost(
	ctx context.Context,
	postID uuid.UUID,
	limit uint64,
	offset uint64,
) (*[]dto.ReactionResponse, error) {
	reactionsResponse, err := rs.reactionRepository.GetReactionsPost(ctx, postID, limit, offset)
	if err != nil {
		return nil, err
	}

	reactions := make([]dto.ReactionResponse, 0, limit)
	for _, reaction := range *reactionsResponse {
		reactions = append(reactions, dto.ReactionResponse{
			ID:       reaction.ID,
			AuthorID: reaction.AuthorID,
			PostID:   reaction.PostID,
		})
	}
	return &reactions, nil
}

func (rs *reactionService) DeleteReaction(
	ctx context.Context,
	reactionID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	return rs.reactionRepository.DeleteReaction(ctx, reactionID, postID, authorID)
}

func (rs *reactionService) DeleteReactionsByAuthor(
	ctx context.Context,
	authorID uuid.UUID,
) error {
	return rs.reactionRepository.DeleteReactionsByAuthor(ctx, authorID)
}

func (rs *reactionService) DeleteReactionsByPost(
	ctx context.Context,
	postID uuid.UUID,
) error {
	return rs.reactionRepository.DeleteReactionsByPost(ctx, postID)
}
