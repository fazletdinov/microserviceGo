package route

import (
	"auth/config"
	"auth/internal/api/controller"

	"auth/internal/domain/repository"
	"auth/internal/domain/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUpdateRouter(env *config.Config, db *gorm.DB, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	updateController := controller.UpdateController{
		UpdateService: service.NewUpdateService(userRepository),
	}
	group.PUT("/user/update", updateController.Update)
}
