package controller

import (
	"net/http"

	"auth/config"
	"auth/internal/domain/service"
	"auth/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteController struct {
	DeleteService service.DeleteService
	Env           *config.Config
}

// DeleteUser	godoc
// @Summary		Удаление пользователя
// @Tags        Auth
// @Accept		json
// @Produce     json
// @Success     204  		{object}  	schemas.SuccessResponse
// @Failure		401			{object}	schemas.ErrorResponse
// @Failure		500			{object}	schemas.ErrorResponse
// @Router      /user/delete [delete]
func (dc *DeleteController) Delete(ctx *gin.Context) {
	userID := ctx.GetString("x-user-id")

	user, err := dc.DeleteService.GetUserByID(ctx, uuid.MustParse(userID))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Message: "Пользователь не авторизован"})
		return
	}

	if errDelete := dc.DeleteService.DeleteUser(ctx, user.ID); errDelete != nil {
		ctx.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: errDelete.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, schemas.SuccessResponse{Message: "Пользователь удален"})

}
