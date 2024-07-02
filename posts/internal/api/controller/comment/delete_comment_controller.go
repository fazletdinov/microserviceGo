package comment

import (
	"net/http"
	"posts/config"
	"posts/internal/domain/service/comment"
	"posts/internal/domain/service/post"
	"posts/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteCommentController struct {
	GetPostService       post.GetPostServcie
	DeleteCommentService comment.DeleteCommentService
	Env                  *config.Config
}

// DeleteComment	godoc
// @Summary		Удаление Comment
// @Tags        Comment
// @Accept		json
// @Produce     json
// @Param	    post_id			path		string		          true		"Post ID"
// @Param	    comment_id		path		string		          true		"Comment ID"
// @Success     204  		    {object}  	schemas.SuccessResponse
// @Failure		401			    {object}	schemas.ErrorResponse
// @Failure		500			    {object}	schemas.ErrorResponse
// @Router      /post/{post_id}/comment/{comment_id} 	[delete]
func (dpc *DeleteCommentController) Delete(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	commentID := ctx.Param("comment_id")

	_, err := dpc.GetPostService.GetByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден"})
		return
	}

	_, err = dpc.DeleteCommentService.GetByID(ctx, uuid.MustParse(postID), uuid.MustParse(commentID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Comment не найден"})
		return
	}

	if err = dpc.DeleteCommentService.DeleteComment(ctx, uuid.MustParse(postID), uuid.MustParse(commentID)); err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, schemas.SuccessResponse{Message: "Comment успешно удален"})
}
