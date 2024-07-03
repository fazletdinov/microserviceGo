package controller

import (
	"likes/config"
	"likes/internal/domain/service"
	"likes/internal/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteReactionController struct {
	DeleteReactionService service.DeleteReactionServcie
	Env                   *config.Config
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

	_, err := dpc.DeleteReactionService.GetByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Reaction не найден"})
		return
	}

	if err = dpc.DeleteReactionService.DeleteReaction(ctx, uuid.MustParse(postID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, schemas.SuccessResponse{Message: "Reaction успешно удален"})

}
