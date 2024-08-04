package repository

import (
	"context"
	"likes/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type reactionRepository struct {
	database *gorm.DB
}

func NewReactionRepository(db *gorm.DB) ReactionRepository {
	return &reactionRepository{
		database: db,
	}
}

func (rr *reactionRepository) Create(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	reaction := models.Reaction{
		PostID:   postID,
		AuthorID: authorID,
	}
	result := rr.database.WithContext(ctx).Create(&reaction)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return reaction.ID, nil
}

func (rr *reactionRepository) GetByIDReaction(
	ctx context.Context,
	reactionID uuid.UUID,
) (*models.Reaction, error) {
	var reaction models.Reaction
	result := rr.database.WithContext(ctx).First(&reaction, "id = ?", reactionID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reaction, nil
}

func (rr *reactionRepository) GetReactionsPost(
	ctx context.Context,
	postID uuid.UUID,
	limit uint64,
	offset uint64,
) (*[]models.Reaction, error) {
	var reactions []models.Reaction
	result := rr.database.WithContext(ctx).Where(&models.Reaction{PostID: postID}).Limit(int(limit)).Offset(int(offset)).Find(&reactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reactions, nil
}

func (rr *reactionRepository) DeleteReaction(
	ctx context.Context,
	reactionID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	result := rr.database.WithContext(ctx).Where("post_id = ? AND author_id = ?", postID, authorID).Delete(&models.Reaction{}, reactionID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (rr *reactionRepository) DeleteReactionsByAuthor(
	ctx context.Context,
	authorID uuid.UUID,
) error {
	result := rr.database.WithContext(ctx).Where("author_id = ?", authorID).Delete(&models.Reaction{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (rr *reactionRepository) DeleteReactionsByPost(
	ctx context.Context,
	postID uuid.UUID,
) error {
	result := rr.database.WithContext(ctx).Where("post_id = ?", postID).Delete(&models.Reaction{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
