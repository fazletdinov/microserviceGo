package post

import (
	"context"
	"posts/internal/models"

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

func (pr *postRepository) Create(
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
	result := pr.database.WithContext(ctx).Create(&post)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return post.ID, nil
}

func (pr *postRepository) GetByIDPost(
	ctx context.Context,
	postID uuid.UUID,
) (*models.Post, error) {
	var post models.Post
	result := pr.database.WithContext(ctx).Preload("Comments").First(&post, "id = ?", postID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func (pr *postRepository) GetPostByIDAuthorID(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*models.Post, error) {
	var post models.Post
	result := pr.database.WithContext(ctx).Where("author_id = ?", authorID).Preload("Comments").First(&post, "id = ?", postID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func (pr *postRepository) GetPosts(
	ctx context.Context,
	limit int,
	offset int,
) (*[]models.Post, error) {
	var posts []models.Post
	result := pr.database.WithContext(ctx).Model(&models.Post{}).Preload("Comments").Limit(int(limit)).Offset(int(offset)).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return &posts, nil
}

func (pr *postRepository) UpdatePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
	title string,
	content string,
) error {
	result := pr.database.WithContext(ctx).Model(&models.Post{ID: postID}).Where("author_id = ?", authorID).Updates(models.Post{Title: title, Content: content})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *postRepository) DeletePost(
	ctx context.Context,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	result := pr.database.WithContext(ctx).Where("author_id = ?", authorID).Delete(&models.Post{}, postID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *postRepository) DeletePostsByAuthor(
	ctx context.Context,
	authorID uuid.UUID,
) error {
	result := pr.database.WithContext(ctx).Where("author_id = ?", authorID).Delete(&models.Post{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
