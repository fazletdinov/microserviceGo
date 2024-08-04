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

type GetCommentController struct {
	GRPCClientPosts *posts.GRPCClientPosts
	Env             *config.Config
}

// ListComments godoc
//
// @Summary		Получение списока Comment
// @Tags	    Comment
// @Accept	    json
// @Produce		json
// @Param	    post_id		    path		       string   true    "Post ID"
// @Param	    limit			query				int		true	"limit"
// @Param	    offset			query				int		true	"offset"
// @Success		200	{array}		schemas.CommentResponse
// @Failure		400	{object}	schemas.ErrorResponse
// @Failure		404	{object}	schemas.ErrorResponse
// @Failure		500	{object}	schemas.ErrorResponse
// @Router	    /post/{post_id}/comments 			[get]
func (gcc *GetCommentController) Fetchs(ctx *gin.Context) {
	var tracer = otel.Tracer(gcc.Env.Jaeger.Application)
	var meter = otel.Meter(gcc.Env.Jaeger.Application)

	postID := ctx.Param("post_id")

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Fetchs",
		oteltrace.WithAttributes(attribute.String("PostID", postID)),
	)
	defer span.End()

	counter, _ := meter.Int64Counter(
		"FetchsComments_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	_, err := gcc.GRPCClientPosts.GetPostByID(traceCtx, uuid.MustParse(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден"})
		return
	}
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

	comments, err := gcc.GRPCClientPosts.GetPostComments(
		traceCtx,
		uuid.MustParse(postID),
		uint64(limit),
		uint64(offset),
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
