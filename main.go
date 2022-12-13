package main

import (
	"github.com/JuliaKozachuk/ProjectStart/ProjectStart/controllers"
	"github.com/JuliaKozachuk/ProjectStart/ProjectStart/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default() //Внутри этой функции инициализируем новый маршрутизатор Gin внутри переменной route.

	models.ConnectDB()

	//route.GET("/", func(context *gin.Context) { //Далее определим GET маршрут до пустого endpoint.
	//context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})

	//})
	route.GET("/tracks", controllers.GetAllTracks)
	route.POST("/tracks", controllers.CreateTrack) // new
	route.GET("/tracks/:id", controllers.GetTrack)
	route.PATCH("/tracks/:id", controllers.UpdateTrack)
	route.DELETE("/tracks/:id", controllers.DeleteTrack)

	route.Run()
}
