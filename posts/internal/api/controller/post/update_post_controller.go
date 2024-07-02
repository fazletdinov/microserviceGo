package post

import (
	"net/http"
	"posts/config"
	"posts/internal/domain/service/post"
	"posts/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdatePostController struct {
	UpdatePostService post.UpdatePostServcie
	Env               *config.Config
}

// UpdatePost   godoc
// @Summary     Обновление Post
// @Tags        Posts
// @Accept		json
// @Produce     json
// @Param	    post_id			path		string					    true    "Post ID"
// @Param		body		    body		schemas.PostUpdateRequest			true	"Для обновления Post"
// @Success     200  		    {object}  	schemas.PostUpdatedResponse
// @Failure	  	400			    {object}	schemas.ErrorResponse
// @Failure	  	401			    {object}	schemas.ErrorResponse
// @Failure	  	500			    {object}	schemas.ErrorResponse
// @Router      /post/{post_id} [put]
func (upc *UpdatePostController) Update(ctx *gin.Context) {
	postID := ctx.Param("post_id")

	post, err := upc.UpdatePostService.GetByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	var postRequest schemas.PostUpdateRequest

	if err = ctx.ShouldBindJSON(&postRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	postRequest.ID = post.ID
	postRequest.AuthorID = post.AuthorID
	err = upc.UpdatePostService.UpdatePost(ctx, &postRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, schemas.SuccessResponse{Message: "Обновление прошло успешно"})

}
