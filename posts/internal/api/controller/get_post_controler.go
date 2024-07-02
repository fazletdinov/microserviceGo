package controller

import (
	"net/http"
	"posts/config"
	"posts/internal/domain/service"
	"posts/internal/schemas"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetPostController struct {
	GetPostService service.GetPostServcie
	Env            *config.Config
}

// GetPost	   godoc
// @Summary	   Получение Post
// @Tags       Posts
// @Accept	   json
// @Produce    json
// @Param	   id			path		string		          	true		"Post ID"
// @Success    200  		{object}  	schemas.PostResponse
// @Failure	   500			{object}	schemas.ErrorResponse
// @Router     /post/{id}    [get]
func (pc *GetPostController) Fetch(ctx *gin.Context) {
	postID := ctx.Param("id")
	post, err := pc.GetPostService.GetByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

// ListPosts godoc
//
// @Summary		Получение списока Post
// @Tags	    Posts
// @Accept	    json
// @Produce		json
// @Param	    limit			query				int		true	"limit"
// @Param	    offset			query				int		true	"offset"
// @Success		200	{array}		schemas.PostResponse
// @Failure		400	{object}	schemas.ErrorResponse
// @Failure		404	{object}	schemas.ErrorResponse
// @Failure		500	{object}	schemas.ErrorResponse
// @Router	    /posts 			[get]
func (pc *GetPostController) Fetchs(ctx *gin.Context) {
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

	posts, err := pc.GetPostService.GetPosts(ctx, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
