package comment

import (
	"context"
	"posts/internal/domain/repository/comment"
	"posts/internal/dto"

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
) (*[]dto.CommentResponse, error) {
	commentsResponse, err := cs.commentRepository.GetComments(ctx, postID, int(limit), int(offset))
	if err != nil {
		return nil, err
	}
	comments := make([]dto.CommentResponse, 0, limit)
	for _, comment := range *commentsResponse {
		comments = append(comments, dto.CommentResponse{
			ID:        comment.ID,
			Text:      comment.Text,
			AuthorID:  comment.AuthorID,
			CreatedAt: comment.CreatedAt,
		})
	}

	return &comments, nil

}

func (cs *commentService) GetCommentByID(
	ctx context.Context,
	commentID uuid.UUID,
	postID uuid.UUID,
	authorID uuid.UUID,
) (*dto.CommentResponse, error) {
	commentResponse, err := cs.commentRepository.GetByIDComment(ctx, commentID, postID, authorID)
	if err != nil {
		return nil, err
	}
	return &dto.CommentResponse{
		ID:        commentResponse.ID,
		Text:      commentResponse.Text,
		AuthorID:  commentResponse.AuthorID,
		CreatedAt: commentResponse.CreatedAt,
	}, nil
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
