package controller

import (
	"net/http"

	"auth/config"
	"auth/internal/domain/service"
	"auth/internal/models"
	"auth/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupService service.SignupService
	Env           *config.Config
}

// LoginUser	godoc
// @Summary		Регистрация пользователя и получение токенов
// @Tags        Auth
// @Accept		json
// @Produce     json
// @Param		body	    body		schemas.SignupRequest	true	"Для создания пользователя и получения токенов"
// @Success     201  		{object}  	schemas.SignupResponse
// @Failure		400			{object}	schemas.ErrorResponse
// @Failure		409			{object}	schemas.ErrorResponse
// @Failure		500			{object}	schemas.ErrorResponse
// @Router      /user/signup [post]
func (sc *SignupController) Signup(ctx *gin.Context) {
	var request schemas.SignupRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.SignupService.GetUserByEmail(ctx, request.Email)
	if err != nil {
		ctx.JSON(http.StatusConflict, schemas.ErrorResponse{Message: "Пользователь с указанным email уже существует"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := models.Users{
		ID:       uuid.New(),
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupService.Create(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := sc.SignupService.CreateAccessToken(&user, sc.Env.JWTConfig.PathPrivateKey, int(sc.Env.JWTConfig.AccessTokenExp))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignupService.CreateRefreshToken(&user, sc.Env.JWTConfig.PathPrivateKey, int(sc.Env.JWTConfig.RefreshTokenExp))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.SetCookie(sc.Env.JWTConfig.SessionCookieName,
		refreshToken,
		int(sc.Env.JWTConfig.RefreshTokenExp),
		"",
		"",
		false,
		true)

	signupResponse := schemas.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusCreated, signupResponse)
}
