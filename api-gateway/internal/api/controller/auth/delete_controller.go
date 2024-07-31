package controller

import (
	"fmt"
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/auth"
	"api-grpc-gateway/internal/clients/likes"
	"api-grpc-gateway/internal/clients/posts"
	schemas "api-grpc-gateway/internal/schemas/auth"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteController struct {
	GRPCClientAuth  *auth.GRPCClientAuth
	GRPCClientPosts *posts.GRPCClientPosts
	GRPCClientLikes *likes.GRPCClientLikes
	Env             *config.Config
}

// DeleteUser	godoc
// @Summary		Удаление пользователя
// @Tags        Auth
// @Accept		json
// @Produce     json
// @Success     204  		{object}  	schemas.SuccessResponse
// @Failure		401			{object}	schemas.ErrorResponse
// @Failure		500			{object}	schemas.ErrorResponse
// @Router      /user/delete [delete]
func (dc *DeleteController) Delete(ctx *gin.Context) {
	var tracer = otel.Tracer(dc.Env.Jaeger.ServerName)
	userID := ctx.GetString("x-user-id")

	_, err := dc.GRPCClientAuth.GetUserByID(ctx, uuid.MustParse(userID))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"DeleteUser",
		oteltrace.WithAttributes(attribute.String("UserID", userID)),
	)
	defer span.End()

	_, errDeleteUser := dc.GRPCClientAuth.DeleteUser(
		traceCtx,
		uuid.MustParse(userID),
	)
	if errDeleteUser != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: fmt.Sprintf("Internal Server %v", err)})
		return
	}

	_, errDeletePost := dc.GRPCClientPosts.DeletePostsByAuthor(
		traceCtx,
		uuid.MustParse(userID),
	)
	if errDeletePost != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: fmt.Sprintf("Internal Server %v", err)})
		return
	}

	_, errDeleteReaction := dc.GRPCClientLikes.DeleteReactionsByAuthor(
		traceCtx,
		uuid.MustParse(userID),
	)
	if errDeleteReaction != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: fmt.Sprintf("Internal Server %v", err)})
		return
	}

	ctx.JSON(http.StatusNoContent, schemas.SuccessResponse{Message: "Пользователь удален"})

}
