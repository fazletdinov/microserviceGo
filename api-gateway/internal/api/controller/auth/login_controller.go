package controller

import (
	"fmt"
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/auth"
	tokenService "api-grpc-gateway/internal/domain/services/auth"
	schemas "api-grpc-gateway/internal/schemas/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	GRPCClientAuth *auth.GRPCClientAuth
	Env            *config.Config
}

// LoginUser    godoc
// @Summary     Получение access и refresh токена в формате JSON
// @Tags        Auth
// @Accept		json
// @Produce     json
// @Param		body		body		schemas.LoginRequest	true	"Для получения токенов"
// @Success     200  		{object}  	schemas.LoginResponse
// @Failure	  	400			{object}	schemas.ErrorResponse
// @Failure	  	401			{object}	schemas.ErrorResponse
// @Failure	  	404			{object}	schemas.ErrorResponse
// @Failure	  	500			{object}	schemas.ErrorResponse
// @Router      /login [post]
func (lc *LoginController) Login(ctx *gin.Context) {
	var tracer = otel.Tracer(lc.Env.Jaeger.ServerName)
	var request schemas.LoginRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"Login",
		oteltrace.WithAttributes(attribute.String("User", request.Email)),
	)
	defer span.End()

	user, err := lc.GRPCClientAuth.GetUserByEmailIsActive(
		traceCtx,
		request.Email,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: fmt.Sprintf("Пользователь с таким email не обнаружен %v", err)})
		return
	}

	if err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	); err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Неверные учетные данные"})
		return
	}

	accessToken, err := tokenService.GenerateAccessToken(
		uuid.MustParse(user.UserId),
		lc.Env.JWTConfig.PathPrivateKey,
		int(lc.Env.JWTConfig.AccessTokenExp),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: fmt.Sprintf("Internal Server error %v", err)})
		return
	}

	refreshToken, err := tokenService.GenerateRefreshToken(
		uuid.MustParse(user.UserId),
		lc.Env.JWTConfig.PathPrivateKey,
		int(lc.Env.JWTConfig.RefreshTokenExp),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: fmt.Sprintf("Internal Server error %v", err)})
		return
	}
	ctx.SetCookie(lc.Env.JWTConfig.SessionCookieName,
		refreshToken,
		int(lc.Env.JWTConfig.RefreshTokenExp),
		"",
		"",
		false,
		true)

	loginResponse := schemas.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, loginResponse)
}
