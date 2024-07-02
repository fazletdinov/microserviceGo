package comment

import (
	"net/http"
	"posts/config"
	"posts/internal/domain/service/comment"
	"posts/internal/domain/service/post"
	"posts/internal/models"
	"posts/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateCommentController struct {
	CreateCommentService comment.CreateCommentServcie
	GetPostService       post.GetPostServcie
	Env                  *config.Config
}

// CreateComment	godoc
// @Summary		Создание Comment
// @Tags        Comment
// @Accept		json
// @Produce     json
// @Param	    post_id					path		string		          				true    "Post ID"
// @Param		body	    			body		schemas.CommentCreateRequest		true	"Создание Comment"
// @Success     201  					{object}  	schemas.SuccessResponse
// @Failure		400						{object}	schemas.ErrorResponse
// @Failure		500						{object}	schemas.ErrorResponse
// @Router      /post/{post_id}/comment [post]
func (ccc *CreateCommentController) Create(ctx *gin.Context) {
	postID := ctx.Param("post_id")

	_, err := ccc.GetPostService.GetByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	var commentRequest schemas.CommentCreateRequest

	if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	comment := models.Comment{
		Text:     commentRequest.Text,
		AuthorID: uuid.New(),
		PostID:   uuid.MustParse(postID),
	}

	if errCreate := ccc.CreateCommentService.CreateComment(ctx, &comment); errCreate != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: errCreate.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.SuccessResponse{Message: "Комментарий успешно создан"})
}
