package post

import (
	"net/http"
	"posts/config"
	"posts/internal/domain/service/post"
	"posts/internal/models"
	"posts/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePostController struct {
	CreatePosteService post.CreatePostServcie
	Env                *config.Config
}

// CreatePost	godoc
// @Summary		Создание Post
// @Tags        Posts
// @Accept		json
// @Produce     json
// @Param		body	    body		schemas.PostCreateRequest		true	"Создание Post"
// @Success     201  		{object}  	schemas.SuccessResponse
// @Failure		400			{object}	schemas.ErrorResponse
// @Failure		500			{object}	schemas.ErrorResponse
// @Router      /post 	[post]
func (pc *CreatePostController) Create(ctx *gin.Context) {
	var postRequest schemas.PostCreateRequest

	if err := ctx.ShouldBindJSON(&postRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	post := models.Post{
		Title:    postRequest.Title,
		Content:  postRequest.Content,
		AuthorID: uuid.New(),
	}
	if err := pc.CreatePosteService.CreatePost(ctx, &post); err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.SuccessResponse{Message: "Post создан успешно"})
}
