package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "posts/docs"
	route "posts/internal/api/router"
	"posts/internal/app"
	autgrpc "posts/internal/clients/auth/grpc"
)

// @title           Gin Posts Service
// @version         1.0
// @description     API-интерфейс службы управления постами и комментариями в Go с использованием платформы Gin framework.

// @contact.name   Идель Фазлетдинов
// @contact.email  fvi-it@mail.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /api/v1

func main() {

	app := app.App()

	env := app.Env

	db := app.DB

	gin := gin.Default()
	gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authClient, err := autgrpc.NewGRPCClient(env.Clients.Address)
	if err != nil {
		panic("Ошибка подключения к gRPC")
	}
	route.SetupRouter(env, db, gin, authClient)

	gin.Run(":" + env.PostsServer.PostsPort)
}
