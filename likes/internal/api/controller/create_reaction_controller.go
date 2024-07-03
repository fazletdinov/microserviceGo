package controller

import (
	"likes/config"
	"likes/internal/domain/service"
	"likes/internal/models"
	"likes/internal/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateReactionController struct {
	CreateReactionService service.CreateReactionServcie
	Env                   *config.Config
}

// CreateReaction godoc
// @Summary		Создание Reaction
// @Tags        Reaction
// @Accept		json
// @Produce     json
// @Param	    post_id			          path		string		            true		"Post ID"
// @Param		body	    body		  schemas.ReactionCreateRequest		true	"Создание Reaction"
// @Success     201  		{object}  	  schemas.SuccessResponse
// @Failure		400			{object}	  schemas.ErrorResponse
// @Failure		500			{object}	  schemas.ErrorResponse
// @Router      /post/{post_id}/reaction  [post]
func (rc *CreateReactionController) Create(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	var reactionRequest schemas.ReactionCreateRequest

	if err := ctx.ShouldBindJSON(&reactionRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	reaction := models.Reaction{
		PostID:   uuid.MustParse(postID),
		AuthorID: reactionRequest.AuthorID,
	}
	if err := rc.CreateReactionService.CreateReaction(ctx, &reaction); err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.SuccessResponse{Message: "Reaction создан успешно"})
}
