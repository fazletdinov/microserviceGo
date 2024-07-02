package route

import (
	"auth/config"
	"auth/internal/api/controller"

	"auth/internal/domain/repository"
	"auth/internal/domain/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProfileRouter(env *config.Config, db *gorm.DB, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	profileController := &controller.ProfileController{
		ProfileService: service.NewProfileService(userRepository),
	}
	group.GET("/user/me", profileController.Fetch)
}
