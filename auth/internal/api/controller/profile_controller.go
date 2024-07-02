package controller

import (
	"net/http"

	"auth/internal/domain/service"
	"auth/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProfileController struct {
	ProfileService service.ProfileService
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

	profile, err := pc.ProfileService.GetProfileByID(ctx, uuid.MustParse(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Пользователь не найден"})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}
