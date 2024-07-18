package controller

import (
	"fmt"
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/likes"
	schemas "api-grpc-gateway/internal/schemas/likes"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateReactionController struct {
	GRPCClientLikes *likes.GRPCClientLikes
	Env             *config.Config
}

// CreateReaction godoc
// @Summary		Создание Reaction
// @Tags        Reaction
// @Accept		json
// @Produce     json
// @Param	    post_id			          path		string		            true		"Post ID"
// @Success     201  		{object}  	  schemas.SuccessResponse
// @Failure		400			{object}	  schemas.ErrorResponse
// @Failure		500			{object}	  schemas.ErrorResponse
// @Router      /post/{post_id}/reaction  [post]
func (rc *CreateReactionController) Create(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	authorID := ctx.GetString("x-user-id")

	reactionID, err := rc.GRPCClientLikes.CreateReaction(ctx, uuid.MustParse(postID), uuid.MustParse(authorID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.SuccessResponse{Message: fmt.Sprintf("ID = %v", reactionID)})
}
