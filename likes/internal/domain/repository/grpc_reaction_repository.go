package repository

import (
	"context"
	"likes/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type reactionGRPCRepository struct {
	database *gorm.DB
}

func NewReactionGRPCRepository(db *gorm.DB) ReactionGRPCRepository {
	return &reactionGRPCRepository{
		database: db,
	}
}

func (rr *reactionGRPCRepository) Create(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	reaction := models.Reaction{
		PostID:   postID,
		AuthorID: authorID,
	}
	result := rr.database.Create(&reaction)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return reaction.ID, nil
}

func (rr *reactionGRPCRepository) GetByIDReaction(
	ctx context.Context,
	reactionID uuid.UUID,
) (*models.Reaction, error) {
	var reaction models.Reaction
	result := rr.database.First(&reaction, "id = ?", reactionID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reaction, nil
}

func (rr *reactionGRPCRepository) GetReactionsPost(
	ctx context.Context,
	postID uuid.UUID,
	limit uint64,
	offset uint64,
) (*[]models.Reaction, error) {
	var reactions []models.Reaction
	result := rr.database.Model(&models.Reaction{PostID: postID}).Limit(int(limit)).Offset(int(offset)).Find(&reactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reactions, nil
}

func (rr *reactionGRPCRepository) DeleteReaction(
	ctx context.Context,
	reactionID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	result := rr.database.Where("post_id = ? AND author_id = ?", postID, authorID).Delete(&models.Reaction{}, reactionID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
