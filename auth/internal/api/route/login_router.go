package route

import (
	"auth/config"

	"auth/internal/api/controller"
	"auth/internal/domain/repository"
	"auth/internal/domain/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewLoginRouter(env *config.Config, db *gorm.DB, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	loginController := &controller.LoginController{
		LoginService: service.NewLoginService(userRepository),
		Env:          env,
	}
	group.POST("/login", loginController.Login)
}
