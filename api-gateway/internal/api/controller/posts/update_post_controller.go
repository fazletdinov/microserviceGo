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

type UpdatePostController struct {
	GRPCClientPosts *posts.GRPCClientPosts
	Env             *config.Config
}

// UpdatePost   godoc
// @Summary     Обновление Post
// @Tags        Posts
// @Accept		json
// @Produce     json
// @Param	    post_id			path		string					    true    "Post ID"
// @Param		body		    body		schemas.UpdatePostRequest	true	"Для обновления Post"
// @Success     200  		    {object}  	schemas.SuccessResponse
// @Failure	  	400			    {object}	schemas.ErrorResponse
// @Failure	  	401			    {object}	schemas.ErrorResponse
// @Failure	  	500			    {object}	schemas.ErrorResponse
// @Router      /post/{post_id} [put]
func (upc *UpdatePostController) Update(ctx *gin.Context) {
	var tracer = otel.Tracer(upc.Env.Jaeger.Application)
	postID := ctx.Param("post_id")
	authorID := ctx.GetString("x-user-id")

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Update",
		oteltrace.WithAttributes(attribute.String("PostID", postID)),
		oteltrace.WithAttributes(attribute.String("AuthorID", authorID)),
	)
	defer span.End()

	_, err := upc.GRPCClientPosts.GetPostByIDAuthorID(
		traceCtx,
		uuid.MustParse(postID),
		uuid.MustParse(authorID),
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Post не найден или вы не являетесь автором поста"})
		return
	}

	var postRequest schemas.UpdatePostRequest

	if err = ctx.ShouldBindJSON(&postRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные  данные"})
		return
	}
	_, err = upc.GRPCClientPosts.UpdatePost(
		traceCtx,
		uuid.MustParse(postID),
		uuid.MustParse(authorID),
		postRequest.Title,
		postRequest.Content,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusOK, schemas.SuccessResponse{Message: "Обновление прошло успешно"})

}
