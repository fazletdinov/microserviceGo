package comment

import (
	"context"
	"posts/internal/domain/repository/comment"
	"posts/internal/models"

	"github.com/google/uuid"
)

type commentGRPCService struct {
	commentGRPCRepository comment.CommentGRPCRepository
}

func NewCommentGRPCService(commentGRPCRepository comment.CommentGRPCRepository) CommentGRPCServcie {
	return &commentGRPCService{
		commentGRPCRepository: commentGRPCRepository,
	}
}

func (cs *commentGRPCService) CreateComment(
	ctx context.Context,
	text string,
	postID uuid.UUID,
	authorID uuid.UUID,
) (uuid.UUID, error) {
	return cs.commentGRPCRepository.CreateComment(ctx, text, postID, authorID)
}

func (cs *commentGRPCService) GetPostComments(
	ctx context.Context,
	postID uuid.UUID,
	limit uint64,
	offset uint64,
) (*[]models.Comment, error) {
	return cs.commentGRPCRepository.GetComments(ctx, postID, int(limit), int(offset))
}

func (cs *commentGRPCService) GetCommentByID(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*models.Comment, error) {
	return cs.commentGRPCRepository.GetByIDComment(ctx, commentID, postID, authorID)
}

func (cs *commentGRPCService) UpdateComment(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
	text string,
) error {
	return cs.commentGRPCRepository.UpdateComment(ctx, commentID, postID, authorID, text)
}

func (cs *commentGRPCService) DeleteComment(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) error {
	return cs.commentGRPCRepository.DeleteComment(ctx, commentID, postID, authorID)
}
