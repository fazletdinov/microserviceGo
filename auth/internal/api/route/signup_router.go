package route

import (
	"auth/config"
	"auth/internal/api/controller"

	"auth/internal/domain/repository"
	"auth/internal/domain/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSignupRouter(env *config.Config, db *gorm.DB, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	userControler := controller.SignupController{
		SignupService: service.NewSignupService(userRepository),
		Env:           env,
	}
	group.POST("/signup", userControler.Signup)
}
