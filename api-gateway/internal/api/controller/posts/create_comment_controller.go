package controller

import (
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/posts"
	schemas "api-grpc-gateway/internal/schemas/posts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateCommentController struct {
	GRPCClientPosts *posts.GRPCClientPosts
	Env             *config.Config
}

// CreateComment	godoc
// @Summary		Создание Comment
// @Tags        Comment
// @Accept		json
// @Produce     json
// @Param	    post_id					path		string		          				true    "Post ID"
// @Param		body	    			body		schemas.CreateCommentRequest		true	"Создание Comment"
// @Success     201  					{object}  	schemas.SuccessResponse
// @Failure		400						{object}	schemas.ErrorResponse
// @Failure		500						{object}	schemas.ErrorResponse
// @Router      /post/{post_id}/comment [post]
func (ccc *CreateCommentController) Create(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	authorID := ctx.GetString("x-user-id")

	_, err := ccc.GRPCClientPosts.GetPostByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден"})
		return
	}

	var commentRequest schemas.CreateCommentRequest

	if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Не валидные данные"})
		return
	}

	if _, errCreate := ccc.GRPCClientPosts.CreateComment(
		ctx,
		commentRequest.Text,
		uuid.MustParse(postID),
		uuid.MustParse(authorID),
	); errCreate != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.SuccessResponse{Message: "Комментарий успешно создан"})
}
