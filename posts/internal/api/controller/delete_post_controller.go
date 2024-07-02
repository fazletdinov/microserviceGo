package controller

import (
	"net/http"
	"posts/config"
	"posts/internal/domain/service"
	"posts/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeletePostController struct {
	DeletePostService service.DeletePostServcie
	Env               *config.Config
}

// DeletePost	godoc
// @Summary		Удаление Post
// @Tags        Posts
// @Accept		json
// @Produce     json
// @Param	    id			    path		string		          true		"Post ID"
// @Success     204  		    {object}  	schemas.SuccessResponse
// @Failure		401			    {object}	schemas.ErrorResponse
// @Failure		500			    {object}	schemas.ErrorResponse
// @Router      /delete/{id} 	[delete]
func (dpc *DeletePostController) Delete(ctx *gin.Context) {
	postID := ctx.Param("id")

	_, err := dpc.DeletePostService.GetByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	if err = dpc.DeletePostService.DeletePost(ctx, uuid.MustParse(postID)); err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, schemas.SuccessResponse{Message: "Post успешно удален"})

}
