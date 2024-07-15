package controller

import (
	"fmt"
	"net/http"

	"api-grpc-gateway/config"
	"api-grpc-gateway/internal/clients/auth"
	schemas "api-grpc-gateway/internal/schemas/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	GRPCClientAuth *auth.GRPCClientAuth
	Env            *config.Config
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
	var request schemas.SignupUserRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	_, err = sc.GRPCClientAuth.GetUserByEmail(ctx, request.Email)
	if err == nil {
		ctx.JSON(http.StatusConflict, schemas.ErrorResponse{Message: "Пользователь с указанным email уже существует"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	userID, err := sc.GRPCClientAuth.CreateUser(ctx, request.Email, string(encryptedPassword))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusCreated, schemas.SuccessResponse{Message: fmt.Sprintf("ID Пользователя = %v", userID)})
}
