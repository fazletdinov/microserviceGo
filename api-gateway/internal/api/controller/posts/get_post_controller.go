package controller

import (
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/posts"
	schemas "api-grpc-gateway/internal/schemas/posts"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var tracer = otel.Tracer(pc.Env.Jaeger.ServerName)
	var meter = otel.Meter(pc.Env.Jaeger.Application)

	postID := ctx.Param("post_id")

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Fetch",
		oteltrace.WithAttributes(attribute.String("PostID", postID)),
	)
	defer span.End()

	counter, _ := meter.Int64Counter(
		"FetchPost_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	post, err := pc.GRPCClientPosts.GetPostByID(
		traceCtx,
		uuid.MustParse(postID),
	)
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
	var tracer = otel.Tracer(pc.Env.Jaeger.Application)
	var meter = otel.Meter(pc.Env.Jaeger.Application)

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

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Fetchs",
		oteltrace.WithAttributes(attribute.Int("Limit", limit)),
		oteltrace.WithAttributes(attribute.Int("Offset", offset)),
	)
	defer span.End()

	counter, _ := meter.Int64Counter(
		"FetchsPosts_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	posts, err := pc.GRPCClientPosts.GetPosts(traceCtx, uint64(limit), uint64(offset))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден"})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
