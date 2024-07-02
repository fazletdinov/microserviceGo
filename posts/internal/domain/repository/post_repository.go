package repository

import (
	"context"
	"posts/internal/models"
	"posts/internal/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postRepository struct {
	database *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		database: db,
	}
}

func (pr *postRepository) Create(ctx context.Context, post *models.Post) error {
	result := pr.database.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *postRepository) GetByIDPost(ctx context.Context, postID uuid.UUID) (*models.Post, error) {
	var post models.Post
	result := pr.database.First(&post, "id = ?", postID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func (pr *postRepository) GetPosts(ctx context.Context, limit int, offset int) (*[]models.Post, error) {
	var posts []models.Post
	result := pr.database.Model(&models.Post{}).Preload("Comments").Limit(int(limit)).Offset(int(offset)).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return &posts, nil
}

func (pr *postRepository) UpdatePost(ctx context.Context, post *schemas.PostUpdateRequest) error {
	result := pr.database.Model(&models.Post{}).Where("id = ?", post.ID).Updates(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *postRepository) DeletePost(ctx context.Context, postID uuid.UUID) error {
	result := pr.database.Delete(&models.Post{}, postID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
