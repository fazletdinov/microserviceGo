package controller

import (
	"fmt"
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/likes"
	schemas "api-grpc-gateway/internal/schemas/likes"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateReactionController struct {
	GRPCClientLikes *likes.GRPCClientLikes
	Env             *config.Config
}

// CreateReaction godoc
// @Summary		Создание Reaction
// @Tags        Reaction
// @Accept		json
// @Produce     json
// @Param	    post_id			          path		string		            true		"Post ID"
// @Success     201  		{object}  	  schemas.SuccessResponse
// @Failure		400			{object}	  schemas.ErrorResponse
// @Failure		500			{object}	  schemas.ErrorResponse
// @Router      /post/{post_id}/reaction  [post]
func (rc *CreateReactionController) Create(ctx *gin.Context) {
	var tracer = otel.Tracer(rc.Env.Jaeger.Application)
	var meter = otel.Meter(rc.Env.Jaeger.Application)

	postID := ctx.Param("post_id")
	authorID := ctx.GetString("x-user-id")

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Create",
		oteltrace.WithAttributes(attribute.String("PostID", postID)),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID)),
	)
	defer span.End()

	counter, _ := meter.Int64Counter(
		"CreateReaction_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	reactionID, err := rc.GRPCClientLikes.CreateReaction(
		traceCtx,
		uuid.MustParse(postID),
		uuid.MustParse(authorID),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.SuccessResponse{Message: fmt.Sprintf("ID = %v", reactionID)})
}
