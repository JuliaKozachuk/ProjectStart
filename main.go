package main

import (
	"net/http"

	"github.com/JuliaKozachuk/ProjectStart/ProjectStart/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default() //Внутри этой функции инициализируем новый маршрутизатор Gin внутри переменной route.

	models.ConnectDB()

	route.GET("/", func(context *gin.Context) { //Далее определим GET маршрут до пустого endpoint.
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})

	})

	route.Run()
}
