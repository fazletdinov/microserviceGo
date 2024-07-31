package controller

import (
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/auth"
	tokenService "api-grpc-gateway/internal/domain/services/auth"
	schemas "api-grpc-gateway/internal/schemas/auth"
	"go.opentelemetry.io/otel"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	GRPCClientAuth *auth.GRPCClientAuth
	Env            *config.Config
}

// Refresh		godoc
// @Summary		Обновление токенов
// @Tags        Auth
// @Accept		json
// @Produce     json
// @Success     200  		{object}  	schemas.RefreshTokenResponse
// @Failure		400			{object}	schemas.ErrorResponse
// @Failure		500			{object}	schemas.ErrorResponse
// @Router      /refresh 	[get]
func (rtc *RefreshTokenController) RefreshToken(ctx *gin.Context) {
	var tracer = otel.Tracer(rtc.Env.Jaeger.ServerName)
	token, errCookie := ctx.Cookie(rtc.Env.JWTConfig.SessionCookieName)
	if errCookie != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	userID, err := tokenService.ExtractUserIDFromToken(
		token,
		rtc.Env.JWTConfig.PathPublicKey,
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	traceCtx, span := tracer.Start(
		ctx.Request.Context(),
		"RefreshToken",
	)
	defer span.End()

	_, err = rtc.GRPCClientAuth.GetUserByID(
		traceCtx,
		userID,
	)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	accessToken, err := tokenService.GenerateAccessToken(
		userID,
		rtc.Env.JWTConfig.PathPrivateKey,
		int(rtc.Env.JWTConfig.AccessTokenExp),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	refreshToken, err := tokenService.GenerateRefreshToken(
		userID,
		rtc.Env.JWTConfig.PathPrivateKey,
		int(rtc.Env.JWTConfig.RefreshTokenExp),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}
	ctx.SetCookie(rtc.Env.JWTConfig.SessionCookieName,
		refreshToken,
		int(rtc.Env.JWTConfig.RefreshTokenExp),
		"",
		"",
		false,
		true)

	refreshTokenResponse := schemas.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, refreshTokenResponse)
}
