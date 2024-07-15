package controller

import (
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/likes"
	schemas "api-grpc-gateway/internal/schemas/likes"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteReactionController struct {
	GRPCClientLikes *likes.GRPCClientLikes
	Env             *config.Config
}

// DeleteReaction	godoc
// @Summary		Удаление Reaction
// @Tags        Reaction
// @Accept		json
// @Produce     json
// @Param	    post_id			         path		string		              true		"Post ID"
// @Success     204  		             {object}  	schemas.SuccessResponse
// @Failure		401			             {object}	schemas.ErrorResponse
// @Failure		500			             {object}	schemas.ErrorResponse
// @Router      /post/{post_id}/reaction [delete]
func (dpc *DeleteReactionController) Delete(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	authorID := ctx.GetString("x-user-id")

	_, err := dpc.GRPCClientLikes.GetReactionByID(ctx, uuid.MustParse(postID), uuid.MustParse(authorID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Reaction не найден"})
		return
	}

	if _, err = dpc.GRPCClientLikes.DeleteReaction(ctx, uuid.MustParse(postID), uuid.MustParse(authorID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusNoContent, schemas.SuccessResponse{Message: "Reaction успешно удален"})

}
