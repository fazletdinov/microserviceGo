package main

import (
	_ "api-grpc-gateway/docs"
	authroute "api-grpc-gateway/internal/api/route/auth"
	likesroute "api-grpc-gateway/internal/api/route/likes"
	postsroute "api-grpc-gateway/internal/api/route/posts"
	authgrpc "api-grpc-gateway/internal/clients/auth"
	likesgrpc "api-grpc-gateway/internal/clients/likes"
	postsgrpc "api-grpc-gateway/internal/clients/posts"
	grpcapp "api-grpc-gateway/internal/grpc_app"
	"api-grpc-gateway/pkg/tracer"
	// "api-grpc-gateway/pkg/logger"
	"api-grpc-gateway/pkg/metric"
	"context"

	"github.com/Cyprinus12138/otelgin"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Gin Api-Gateway Service
// @version         1.0
// @description     API-интерфейс, выступающий в роли шлюза для управления постами, комментариями и пользователями в Go с использованием платформы Gin framework.

// @contact.name   Идель Фазлетдинов
// @contact.email  fvi-it@mail.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

func main() {
	app := grpcapp.App()

	env := app.Env
	// log := app.Log

	ctx := context.Background()
	{
		tp, err := tracer.InitTracer(
			ctx,
			env.Jaeger.CollectorUrl,
			env.Jaeger.Application,
		)
		if err != nil {
			panic(err)
		}
		defer tp.Shutdown(ctx)
		mp, err := metric.SetupMetrics(
			ctx,
			env.Jaeger.Application,
		)
		if err != nil {
			panic(err)
		}
		defer mp.Shutdown(ctx)
	}

	gin := gin.Default()
	// gin.Use(logger.LoggingMiddleware(log, env.Jaeger.Application))
	gin.Use(otelgin.Middleware(env.Jaeger.Application))

	gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authClient, err := authgrpc.NewGRPCClientAuth(
		env.GatewayGRPCServer.AuthServerAddress,
		env,
	)
	if err != nil {
		panic("Ошибка подключения к сервису Auth")
	}
	postsClient, err := postsgrpc.NewGRPCClientPosts(
		env.GatewayGRPCServer.PostsServerAddress,
		env,
	)
	if err != nil {
		panic("Ошибка подключения к сервису Posts")
	}
	likesClient, err := likesgrpc.NewGRPCClientLikes(
		env.GatewayGRPCServer.LikesServerAddress,
		env,
	)
	if err != nil {
		panic("Ошибка подключения к сервису Likes")
	}
	authroute.SetupAuthRouter(gin, authClient, postsClient, likesClient, env)
	postsroute.SetupPostsRouter(gin, postsClient, likesClient, env)
	likesroute.SetupLikesRouter(gin, likesClient, env)

	gin.Run(":" + env.GatewayGRPCServer.ApiGatewayPort)
}
