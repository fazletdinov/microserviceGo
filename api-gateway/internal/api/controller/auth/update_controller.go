package controller

import (
	"net/http"

	"api-grpc-gateway/internal/clients/auth"
	schemas "api-grpc-gateway/internal/schemas/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateController struct {
	GRPCClientAuth *auth.GRPCClientAuth
}

// UpdateUser    godoc
// @Summary     Обновление пользователя
// @Tags        Auth
// @Accept		json
// @Produce     json
// @Param		body		body		schemas.UpdateUser			true	"Для получения обновления пользователя"
// @Success     200  		{object}  	schemas.SuccessResponse
// @Failure	  	400			{object}	schemas.ErrorResponse
// @Failure	  	401			{object}	schemas.ErrorResponse
// @Failure	  	500			{object}	schemas.ErrorResponse
// @Router      /user/update [put]
func (uc *UpdateController) Update(ctx *gin.Context) {
	userID := ctx.GetString("x-user-id")
	_, err := uc.GRPCClientAuth.GetUserByID(ctx, uuid.MustParse(userID))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	var userUpdate schemas.UpdateUser
	if err = ctx.ShouldBindJSON(&userUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	_, err = uc.GRPCClientAuth.UpdateUser(ctx, uuid.MustParse(userID), *userUpdate.FirstName, *userUpdate.LastName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusOK, schemas.SuccessResponse{Message: "Обновление прошло успешно"})
}
