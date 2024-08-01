package controller

import (
	"fmt"
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/auth"
	schemas "api-grpc-gateway/internal/schemas/auth"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProfileController struct {
	GRPCClientAuth *auth.GRPCClientAuth
	Env            *config.Config
}

// GetUser		godoc
// @Summary		Получение пользователя
// @Tags        Auth
// @Accept		json
// @Produce     json
// @Success     200  		{object}  	schemas.UserResponse
// @Failure		401			{object}	schemas.ErrorResponse
// @Failure		404			{object}	schemas.ErrorResponse
// @Router      /user/me     [get]
func (pc *ProfileController) Fetch(ctx *gin.Context) {
	var tracer = otel.Tracer(pc.Env.Jaeger.Application)
	fmt.Printf("tracer ===================== %v\n", tracer)
	userID := ctx.GetString("x-user-id")

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Fetch",
		oteltrace.WithAttributes(attribute.String("UserID", userID)),
	)
	fmt.Printf("tracer ===================== %v\n", traceCtx)
	fmt.Printf("span ===================== %v\n", span)

	profile, err := pc.GRPCClientAuth.GetUserByID(
		traceCtx,
		uuid.MustParse(userID),
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	defer span.End()

	ctx.JSON(http.StatusOK, profile)
}
