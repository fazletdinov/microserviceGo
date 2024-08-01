package controller

import (
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/likes"
	schemas "api-grpc-gateway/internal/schemas/likes"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetReactionController struct {
	GRPCClientLikes *likes.GRPCClientLikes
	Env             *config.Config
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
	var tracer = otel.Tracer(rc.Env.Jaeger.Application)
	postID := ctx.Param("post_id")

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные данные limit"})
		return
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные данные offset"})
		return
	}

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Fetchs",
		oteltrace.WithAttributes(attribute.String("PostID", postID)),
		oteltrace.WithAttributes(attribute.Int("Limit", limit)),
		oteltrace.WithAttributes(attribute.Int("Offset", offset)),
	)
	defer span.End()

	reactions, err := rc.GRPCClientLikes.GetReactionsPost(
		traceCtx,
		uuid.MustParse(postID),
		uint64(limit),
		uint64(offset),
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "not found"})
		return
	}

	ctx.JSON(http.StatusOK, reactions)
}
