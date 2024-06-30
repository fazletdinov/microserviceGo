package route

import (
	"auth/config"
	"auth/internal/api/controller"

	"auth/internal/domain/repository"
	"auth/internal/domain/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRefreshTokenRouter(env *config.Config, db *gorm.DB, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	refreshTokenController := &controller.RefreshTokenController{
		RefreshTokenService: service.NewRefreshTokenService(userRepository),
		Env:                 env,
	}
	group.GET("/refresh", refreshTokenController.RefreshToken)
}
