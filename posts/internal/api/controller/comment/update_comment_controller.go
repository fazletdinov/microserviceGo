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

type UpdateCommentController struct {
	GetPostService       post.GetPostServcie
	UpdateCommentService comment.UpdateCommentService
	Env                  *config.Config
}

// UpdateComment   godoc
// @Summary     Обновление Comment
// @Tags        Comment
// @Accept		json
// @Produce     json
// @Param	    post_id			path		string					      true    "Post ID"
// @Param	    comment_id		path		string					      true    "Comment ID"
// @Param		body		    body		schemas.CommentUpdateRequest  true	"Для обновления Post"
// @Success     200  		    {object}  	schemas.SuccessResponse
// @Failure	  	400			    {object}	schemas.ErrorResponse
// @Failure	  	401			    {object}	schemas.ErrorResponse
// @Failure	  	500			    {object}	schemas.ErrorResponse
// @Router      /post/{post_id}/comment/{comment_id} 	[put]
func (upc *UpdateCommentController) Update(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	commentID := ctx.Param("comment_id")

	_, err := upc.GetPostService.GetByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = upc.UpdateCommentService.GetByID(ctx, uuid.MustParse(postID), uuid.MustParse(commentID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	var commentRequest schemas.CommentUpdateRequest

	if err = ctx.ShouldBindJSON(&commentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	err = upc.UpdateCommentService.UpdateComment(ctx, uuid.MustParse(postID), uuid.MustParse(commentID), &commentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, schemas.SuccessResponse{Message: "Обновление прошло успешно"})

}
