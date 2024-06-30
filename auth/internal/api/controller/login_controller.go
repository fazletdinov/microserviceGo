package controller

import (
	"fmt"
	"net/http"

	"auth/config"
	"auth/internal/schemas"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginService schemas.LoginService
	Env          *config.Config
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
	var request schemas.LoginRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lc.LoginService.GetUserByEmail(ctx, request.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Пользователь с таким email не обнаружен"})
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: fmt.Sprintf("Неверные учетные данные, %v", user.Password)})
		return
	}

	accessToken, err := lc.LoginService.CreateAccessToken(&user, lc.Env.JWTConfig.PathPrivateKey, int(lc.Env.JWTConfig.AccessTokenExp))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginService.CreateRefreshToken(&user, lc.Env.JWTConfig.PathPrivateKey, int(lc.Env.JWTConfig.RefreshTokenExp))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
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
