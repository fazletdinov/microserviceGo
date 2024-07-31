package controller

import (
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/posts"
	schemas "api-grpc-gateway/internal/schemas/posts"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteCommentController struct {
	GRPCClientPosts *posts.GRPCClientPosts
	Env             *config.Config
}

// DeleteComment	godoc
// @Summary		Удаление Comment
// @Tags        Comment
// @Accept		json
// @Produce     json
// @Param	    post_id			path		string		          true		"Post ID"
// @Param	    comment_id		path		string		          true		"Comment ID"
// @Success     204  		    {object}  	schemas.SuccessResponse
// @Failure		401			    {object}	schemas.ErrorResponse
// @Failure		500			    {object}	schemas.ErrorResponse
// @Router      /post/{post_id}/comment/{comment_id} 	[delete]
func (dpc *DeleteCommentController) Delete(ctx *gin.Context) {
	var tracer = otel.Tracer(dpc.Env.Jaeger.ServerName)
	postID := ctx.Param("post_id")
	commentID := ctx.Param("comment_id")
	authorID := ctx.GetString("x-user-id")

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Delete",
		oteltrace.WithAttributes(attribute.String("PostID", postID)),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID)),
		oteltrace.WithAttributes(attribute.String("CommentID", commentID)),
	)
	defer span.End()

	_, err := dpc.GRPCClientPosts.GetPostByID(
		traceCtx,
		uuid.MustParse(postID),
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден"})
		return
	}

	_, err = dpc.GRPCClientPosts.GetCommentByID(
		traceCtx,
		uuid.MustParse(commentID),
		uuid.MustParse(postID),
		uuid.MustParse(authorID),
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Comment не найден"})
		return
	}

	if _, err = dpc.GRPCClientPosts.DeleteComment(
		traceCtx,
		uuid.MustParse(commentID),
		uuid.MustParse(postID),
		uuid.MustParse(authorID),
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusNoContent, schemas.SuccessResponse{Message: "Comment успешно удален"})
}
