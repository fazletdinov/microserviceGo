package controller

import (
	"fmt"
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/posts"
	schemas "api-grpc-gateway/internal/schemas/posts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateCommentController struct {
	GRPCClientPosts *posts.GRPCClientPosts
	Env             *config.Config
}

// UpdateComment   godoc
// @Summary     Обновление Comment
// @Tags        Comment
// @Accept		json
// @Produce     json
// @Param	    post_id			path		string					      true    "Post ID"
// @Param	    comment_id		path		string					      true    "Comment ID"
// @Param		body		    body		schemas.UpdateCommentRequest  true	"Для обновления Post"
// @Success     200  		    {object}  	schemas.SuccessResponse
// @Failure	  	400			    {object}	schemas.ErrorResponse
// @Failure	  	401			    {object}	schemas.ErrorResponse
// @Failure	  	500			    {object}	schemas.ErrorResponse
// @Router      /post/{post_id}/comment/{comment_id} 	[put]
func (upc *UpdateCommentController) Update(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	commentID := ctx.Param("comment_id")
	authorID := ctx.GetString("x-user-id")

	_, err := upc.GRPCClientPosts.GetPostByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден"})
		return
	}

	_, err = upc.GRPCClientPosts.GetCommentByID(ctx, uuid.MustParse(commentID), uuid.MustParse(postID), uuid.MustParse(authorID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Comment не найден"})
		return
	}

	var commentRequest schemas.UpdateCommentRequest

	if err = ctx.ShouldBindJSON(&commentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные данные"})
		return
	}
	_, err = upc.GRPCClientPosts.UpdateComment(ctx, uuid.MustParse(commentID), commentRequest.Text, uuid.MustParse(postID), uuid.MustParse(authorID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusOK, schemas.SuccessResponse{Message: "Обновление прошло успешно"})

}
