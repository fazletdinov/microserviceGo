package controller

import (
	"fmt"
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/likes"
	"api-grpc-gateway/internal/clients/posts"
	schemas "api-grpc-gateway/internal/schemas/posts"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeletePostController struct {
	GRPCClientPosts *posts.GRPCClientPosts
	GRPCClientLikes *likes.GRPCClientLikes
	Env             *config.Config
}

// DeletePost	godoc
// @Summary		Удаление Post
// @Tags        Posts
// @Accept		json
// @Produce     json
// @Param	    post_id			path		string		          		true 	"Post ID"
// @Success     204  		    {object}  	schemas.SuccessResponse
// @Failure		401			    {object}	schemas.ErrorResponse
// @Failure		500			    {object}	schemas.ErrorResponse
// @Router      /post/{post_id} [delete]
func (dpc *DeletePostController) Delete(ctx *gin.Context) {
	var tracer = otel.Tracer(dpc.Env.Jaeger.Application)
	var meter = otel.Meter(dpc.Env.Jaeger.Application)

	postID := ctx.Param("post_id")
	authorID := ctx.GetString("x-user-id")

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Delete",
		oteltrace.WithAttributes(attribute.String("PostID", postID)),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID)),
	)
	defer span.End()

	counter, _ := meter.Int64Counter(
		"DeletePost_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	_, err := dpc.GRPCClientPosts.GetPostByIDAuthorID(
		traceCtx,
		uuid.MustParse(postID),
		uuid.MustParse(authorID),
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: fmt.Sprintf("Post не найден или вы не являетесь автором поста %v", err)})
		return
	}

	if _, err = dpc.GRPCClientPosts.DeletePost(
		traceCtx,
		uuid.MustParse(postID),
		uuid.MustParse(authorID),
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	if _, err = dpc.GRPCClientLikes.DeleteReactionsByPost(
		traceCtx,
		uuid.MustParse(postID),
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusNoContent, schemas.SuccessResponse{Message: "Post успешно удален"})

}
