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

type CreatePostController struct {
	GRPCClientPosts *posts.GRPCClientPosts
	Env             *config.Config
}

// CreatePost	godoc
// @Summary		Создание Post
// @Tags        Posts
// @Accept		json
// @Produce     json
// @Param		body	    body		schemas.CreatePostRequest		true	"Создание Post"
// @Success     201  		{object}  	schemas.SuccessResponse
// @Failure		400			{object}	schemas.ErrorResponse
// @Failure		500			{object}	schemas.ErrorResponse
// @Router      /post 		[post]
func (cp *CreatePostController) Create(ctx *gin.Context) {
	var request schemas.CreatePostRequest
	userID := ctx.GetString("x-user-id")

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	postID, err := cp.GRPCClientPosts.CreatePost(ctx, request.Title, request.Content, uuid.MustParse(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.SuccessResponse{Message: fmt.Sprintf("ID = %v", postID)})
}
