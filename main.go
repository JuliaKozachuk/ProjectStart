package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JuliaKozachuk/ProjectStart/controllers"
	"github.com/JuliaKozachuk/ProjectStart/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	//"host=127.0.0.1 port=5432 user=postgres dbname=musicstore password=qwerty sslmode=disable")
	postgres_data := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s ", host, port, user, dbname, password, sslmode)
	route := gin.Default() //Внутри этой функции инициализируем новый маршрутизатор Gin внутри переменной route.

	models.ConnectDB(postgres_data)

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
