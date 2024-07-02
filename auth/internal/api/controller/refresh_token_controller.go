package controller

import (
	"net/http"

	"auth/config"
	"auth/internal/domain/service"
	"auth/internal/schemas"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	RefreshTokenService service.RefreshTokenService
	Env                 *config.Config
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
	token, errCookie := ctx.Cookie(rtc.Env.JWTConfig.SessionCookieName)
	if errCookie != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	id, err := rtc.RefreshTokenService.ExtractIDFromToken(token, rtc.Env.JWTConfig.PathPublicKey)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	user, err := rtc.RefreshTokenService.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	accessToken, err := rtc.RefreshTokenService.CreateAccessToken(user, rtc.Env.JWTConfig.PathPrivateKey, int(rtc.Env.JWTConfig.AccessTokenExp))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := rtc.RefreshTokenService.CreateRefreshToken(user, rtc.Env.JWTConfig.PathPrivateKey, int(rtc.Env.JWTConfig.RefreshTokenExp))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
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
