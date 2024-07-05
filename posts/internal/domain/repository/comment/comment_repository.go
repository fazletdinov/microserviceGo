package comment

import (
	"context"
	"posts/internal/models"
	"posts/internal/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type commentRepository struct {
	database *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		database: db,
	}
}

func (cr *commentRepository) CreateComment(ctx context.Context, comment *models.Comment) error {
	result := cr.database.Create(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *commentRepository) GetByIDComment(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID) (*models.Comment, error) {
	var comment models.Comment
	result := cr.database.Model(&models.Comment{PostID: postID, AuthorID: authorID}).First(&comment, "id = ?", commentID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func (cr *commentRepository) GetComments(ctx context.Context, postID uuid.UUID, limit int, offset int) (*[]models.Comment, error) {
	var comments []models.Comment
	result := cr.database.Model(&models.Comment{PostID: postID}).Limit(int(limit)).Offset(int(offset)).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comments, nil
}

func (cr *commentRepository) UpdateComment(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID, comment *schemas.CommentUpdateRequest) error {
	result := cr.database.Model(&models.Comment{PostID: postID, ID: commentID, AuthorID: authorID}).Updates(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *commentRepository) DeleteComment(ctx context.Context, postID uuid.UUID, commentID uuid.UUID, authorID uuid.UUID) error {
	result := cr.database.Where("post_id = ? AND author_id = ?", postID, authorID).Delete(&models.Comment{}, commentID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
