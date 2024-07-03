package controller

import (
	"likes/config"
	"likes/internal/domain/service"
	"likes/internal/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetReactionController struct {
	GetReactionService service.GetReactionServcie
	Env                *config.Config
}

// ListReactions godoc
//
// @Summary		Получение списока Reactin на Post
// @Tags	    Reaction
// @Accept	    json
// @Produce		json
// @Param	    post_id			path		      string    true    "Post ID"
// @Param	    limit			query				int		true	"limit"
// @Param	    offset			query				int		true	"offset"
// @Success		200	{array}		schemas.ReactionResponse
// @Failure		400	{object}	schemas.ErrorResponse
// @Failure		404	{object}	schemas.ErrorResponse
// @Failure		500	{object}	schemas.ErrorResponse
// @Router	    /post/{post_id}/reactions 	[get]
func (rc *GetReactionController) Fetchs(ctx *gin.Context) {
	postID := ctx.Param("post_id")

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

	reactions, err := rc.GetReactionService.GetReactionsPost(ctx, uuid.MustParse(postID), limit, offset)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "not found"})
		return
	}

	ctx.JSON(http.StatusOK, reactions)
}
