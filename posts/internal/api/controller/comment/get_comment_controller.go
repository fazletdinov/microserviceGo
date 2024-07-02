package comment

import (
	"net/http"
	"posts/config"
	"posts/internal/domain/service/comment"
	"posts/internal/domain/service/post"
	"posts/internal/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetCommentController struct {
	GetCommentService comment.GetCommentServcie
	GetPostService    post.GetPostServcie
	Env               *config.Config
}

// GetComment  godoc
// @Summary	   Получение Comment
// @Tags       Comment
// @Accept	   json
// @Produce    json
// @Param	   post_id		path		string		          	true		"Post ID"
// @Param	   comment_id	path		string		          	true		"Comment ID"
// @Success    200  		{object}  	schemas.CommentResponse
// @Failure	   500			{object}	schemas.ErrorResponse
// @Router     /post/{post_id}/comment/{comment_id}  [get]
func (gcc *GetCommentController) Fetch(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	commentID := ctx.Param("comment_id")
	_, err := gcc.GetPostService.GetByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден"})
		return
	}

	comment, err := gcc.GetCommentService.GetByID(ctx, uuid.MustParse(postID), uuid.MustParse(commentID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Comment не найден"})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// ListComments godoc
//
// @Summary		Получение списока Comment
// @Tags	    Comment
// @Accept	    json
// @Produce		json
// @Param	    post_id		    path		       string   true    "Post ID"
// @Param	    limit			query				int		true	"limit"
// @Param	    offset			query				int		true	"offset"
// @Success		200	{array}		schemas.CommentResponse
// @Failure		400	{object}	schemas.ErrorResponse
// @Failure		404	{object}	schemas.ErrorResponse
// @Failure		500	{object}	schemas.ErrorResponse
// @Router	    /post/{post_id}/comments 			[get]
func (gcc *GetCommentController) Fetchs(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	_, err := gcc.GetPostService.GetByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден"})
		return
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	comments, err := gcc.GetCommentService.GetComments(ctx, uuid.MustParse(postID), limit, offset)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
