package main

import (
	_ "auth/docs"
	route "auth/internal/api/route"
	"auth/internal/app"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Gin Auth Service
// @version         1.0
// @description     API-интерфейс службы управления пользователями в Go с использованием платформы Gin framework.

// @contact.name   Идель Фазлетдинов
// @contact.email  fvi-it@mail.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {

	app := app.App()

	env := app.Env

	db := app.DB

	gin := gin.Default()
	gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	route.SetupRouter(env, db, gin)

	gin.Run(":" + env.AuthServer.AuthPort)
}
