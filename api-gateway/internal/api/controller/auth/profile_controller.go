package controller

import (
	"net/http"

	"api-grpc-gateway/internal/clients/auth"
	schemas "api-grpc-gateway/internal/schemas/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProfileController struct {
	GRPCClientAuth *auth.GRPCClientAuth
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
	userID := ctx.GetString("x-user-id")

	profile, err := pc.GRPCClientAuth.GetUserByID(ctx, uuid.MustParse(userID))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}
