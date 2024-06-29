package main

import (
	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	// Создайте новый роутер по умолчанию у gin
	router := gin.Default()
	// Регистрируем обработчик GET-запроса /ping
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Привет я Gin",
		})
	})
	router.Run(":9100")

}
