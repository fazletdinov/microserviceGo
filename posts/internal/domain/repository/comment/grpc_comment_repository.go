package comment

import (
	"context"
	"posts/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type commentGRPCRepository struct {
	database *gorm.DB
}

func NewCommentGRPCRepository(db *gorm.DB) CommentGRPCRepository {
	return &commentGRPCRepository{
		database: db,
	}
}

func (cr *commentGRPCRepository) CreateComment(
	ctx context.Context,
	text string,
	postID uuid.UUID,
	authorID uuid.UUID) (uuid.UUID, error) {
	comment := models.Comment{
		Text:     text,
		PostID:   postID,
		AuthorID: authorID,
	}
	result := cr.database.WithContext(ctx).Create(&comment)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return comment.ID, nil
}

func (cr *commentGRPCRepository) GetByIDComment(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*models.Comment, error) {
	var comment models.Comment
	result := cr.database.WithContext(ctx).Model(&models.Comment{PostID: postID, AuthorID: authorID}).First(&comment, "id = ?", commentID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func (cr *commentGRPCRepository) GetComments(
	ctx context.Context,
	postID uuid.UUID,
	limit int,
	offset int,
) (*[]models.Comment, error) {
	var comments []models.Comment
	result := cr.database.WithContext(ctx).Model(&models.Comment{PostID: postID}).Limit(limit).Offset(offset).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comments, nil
}

func (cr *commentGRPCRepository) UpdateComment(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
	text string,
) error {
	result := cr.database.WithContext(ctx).Model(&models.Comment{PostID: postID, ID: commentID, AuthorID: authorID}).Updates(models.Comment{Text: text})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *commentGRPCRepository) DeleteComment(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	result := cr.database.WithContext(ctx).Where("post_id = ? AND author_id = ?", postID, authorID).Delete(&models.Comment{}, commentID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
