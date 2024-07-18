package controller

import (
	"fmt"
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/likes"
	"api-grpc-gateway/internal/clients/posts"
	schemas "api-grpc-gateway/internal/schemas/posts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeletePostController struct {
	GRPCClientPosts *posts.GRPCClientPosts
	GRPCClientLikes *likes.GRPCClientLikes
	Env             *config.Config
}

// DeletePost	godoc
// @Summary		Удаление Post
// @Tags        Posts
// @Accept		json
// @Produce     json
// @Param	    post_id			path		string		          		true 	"Post ID"
// @Success     204  		    {object}  	schemas.SuccessResponse
// @Failure		401			    {object}	schemas.ErrorResponse
// @Failure		500			    {object}	schemas.ErrorResponse
// @Router      /post/{post_id} [delete]
func (dpc *DeletePostController) Delete(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	authorID := ctx.GetString("x-user-id")

	_, err := dpc.GRPCClientPosts.GetPostByIDAuthorID(ctx, uuid.MustParse(postID), uuid.MustParse(authorID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: fmt.Sprintf("Post не найден или вы не являетесь автором поста %v", err)})
		return
	}

	if _, err = dpc.GRPCClientPosts.DeletePost(ctx, uuid.MustParse(postID), uuid.MustParse(authorID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	if _, err = dpc.GRPCClientLikes.DeleteReactionsByPost(ctx, uuid.MustParse(postID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusNoContent, schemas.SuccessResponse{Message: "Post успешно удален"})

}
