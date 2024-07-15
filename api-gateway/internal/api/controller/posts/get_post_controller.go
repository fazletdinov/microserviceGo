package controller

import (
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/posts"
	schemas "api-grpc-gateway/internal/schemas/posts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
)

type GetPostController struct {
	GRPCClientPosts *posts.GRPCClientPosts
	Env             *config.Config
}

// GetPost	   godoc
// @Summary	   Получение Post
// @Tags       Posts
// @Accept	   json
// @Produce    json
// @Param	   post_id		      path		    string		          	true		"Post ID"
// @Success    200  		      {object}  	schemas.PostResponse
// @Failure	   500			      {object}	    schemas.ErrorResponse
// @Router     /post/{post_id}    [get]
func (pc *GetPostController) Fetch(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	post, err := pc.GRPCClientPosts.GetPostByID(ctx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден"})
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
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные данные"})
		return
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	posts, err := pc.GRPCClientPosts.GetPosts(ctx, uint64(limit), uint64(offset))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден"})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
