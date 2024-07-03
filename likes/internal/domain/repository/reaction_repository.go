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

func (rr *reactionRepository) Create(ctx context.Context, reaction *models.Reaction) error {
	result := rr.database.Create(&reaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (rr *reactionRepository) GetByIDReaction(ctx context.Context, reactionID uuid.UUID) (*models.Reaction, error) {
	var reaction models.Reaction
	result := rr.database.First(&reaction, "id = ?", reactionID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reaction, nil
}

func (rr *reactionRepository) GetReactionsPost(ctx context.Context, postID uuid.UUID, limit int, offset int) (*[]models.Reaction, error) {
	var reactions []models.Reaction
	result := rr.database.Model(&models.Reaction{PostID: postID}).Limit(int(limit)).Offset(int(offset)).Find(&reactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reactions, nil
}

func (rr *reactionRepository) DeleteReaction(ctx context.Context, reactionID uuid.UUID) error {
	result := rr.database.Delete(&models.Reaction{}, reactionID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
