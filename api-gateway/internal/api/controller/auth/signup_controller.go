package controller

import (
	"fmt"
	"log/slog"
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/auth"
	schemas "api-grpc-gateway/internal/schemas/auth"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	GRPCClientAuth *auth.GRPCClientAuth
	Env            *config.Config
	Log            *slog.Logger
}

// LoginUser	godoc
// @Summary		Регистрация пользователя и получение токенов
// @Tags        Auth
// @Accept		json
// @Produce     json
// @Param		body	    body		schemas.SignupUserRequest	true	"Для создания пользователя и получения токенов"
// @Success     201  		{object}  	schemas.SuccessResponse
// @Failure		400			{object}	schemas.ErrorResponse
// @Failure		409			{object}	schemas.ErrorResponse
// @Failure		500			{object}	schemas.ErrorResponse
// @Router      /user/signup [post]
func (sc *SignupController) Signup(ctx *gin.Context) {
	var tracer = otel.Tracer(sc.Env.Jaeger.Application)
	var meter = otel.Meter(sc.Env.Jaeger.Application)

	var request schemas.SignupUserRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Signup",
		oteltrace.WithAttributes(attribute.String("Email", request.Email)),
	)
	defer span.End()

	counter, _ := meter.Int64Counter(
		"Signup_counter",
	)
	counter.Add(
		ctx,
		1,
	)

	_, err = sc.GRPCClientAuth.GetUserByEmail(
		traceCtx,
		request.Email,
	)
	if err == nil {
		ctx.JSON(http.StatusConflict, schemas.ErrorResponse{Message: "Пользователь с указанным email уже существует"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: fmt.Sprintf("Internal Server error %v", err)})
		return
	}

	userID, err := sc.GRPCClientAuth.CreateUser(
		traceCtx,
		request.Email,
		string(encryptedPassword),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: fmt.Sprintf("Internal Server error %v", err)})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.SuccessResponse{Message: fmt.Sprintf("ID Пользователя = %v", userID)})
}
