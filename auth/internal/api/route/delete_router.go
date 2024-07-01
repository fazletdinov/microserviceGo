package route

import (
	"auth/config"
	"auth/internal/api/controller"

	"auth/internal/domain/repository"
	"auth/internal/domain/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewDeleteRouter(env *config.Config, db *gorm.DB, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	deleteController := controller.DeleteController{
		DeleteService: service.NewDeleteService(userRepository),
	}
	group.DELETE("/delete", deleteController.Delete)
}
