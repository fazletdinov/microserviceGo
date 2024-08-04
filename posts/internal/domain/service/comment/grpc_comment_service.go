package comment

import (
	"context"
	"posts/internal/domain/repository/comment"
	"posts/internal/models"

	"github.com/google/uuid"
)

type commentService struct {
	commentRepository comment.CommentRepository
}

func NewCommentService(commentRepository comment.CommentRepository) CommentServcie {
	return &commentService{
		commentRepository: commentRepository,
	}
}

func (cs *commentService) CreateComment(
	ctx context.Context,
	text string,
	postID uuid.UUID,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	return cs.commentRepository.CreateComment(ctx, text, postID, authorID)
}

func (cs *commentService) GetPostComments(
	ctx context.Context,
	postID uuid.UUID,
	limit uint64,
	offset uint64,
) (*[]models.Comment, error) {
	return cs.commentRepository.GetComments(ctx, postID, int(limit), int(offset))
}

func (cs *commentService) GetCommentByID(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*models.Comment, error) {
	return cs.commentRepository.GetByIDComment(ctx, commentID, postID, authorID)
}

func (cs *commentService) UpdateComment(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
	text string,
) error {
	return cs.commentRepository.UpdateComment(ctx, commentID, postID, authorID, text)
}

func (cs *commentService) DeleteComment(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	return cs.commentRepository.DeleteComment(ctx, commentID, postID, authorID)
}
