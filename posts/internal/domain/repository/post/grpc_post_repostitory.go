package post

import (
	"context"
	"posts/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postGRPCRepository struct {
	database *gorm.DB
}

func NewPostGRPCRepository(db *gorm.DB) PostGRPCRepository {
	return &postGRPCRepository{
		database: db,
	}
}

func (pr *postGRPCRepository) Create(
	ctx context.Context,
	title string,
	content string,
	authirID uuid.UUID,
) (uuid.UUID, error) {
	post := models.Post{
		Title:    title,
		Content:  content,
		AuthorID: authirID,
	}
	result := pr.database.Create(&post)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return post.ID, nil
}

func (pr *postGRPCRepository) GetByIDPost(
	ctx context.Context,
	postID uuid.UUID,
) (*models.Post, error) {
	var post models.Post
	result := pr.database.Preload("Comments").First(&post, "id = ?", postID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func (pr *postGRPCRepository) GetPosts(
	ctx context.Context,
	limit int,
	offset int,
) (*[]models.Post, error) {
	var posts []models.Post
	result := pr.database.Model(&models.Post{}).Preload("Comments").Limit(int(limit)).Offset(int(offset)).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return &posts, nil
}

func (pr *postGRPCRepository) UpdatePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
	title string,
	content string,
) error {
	result := pr.database.Model(&models.Post{ID: postID}).Where("author_id = ?", authorID).Updates(models.Post{Title: title, Content: content})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *postGRPCRepository) DeletePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	result := pr.database.Where("author_id = ?", authorID).Delete(&models.Post{}, postID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
