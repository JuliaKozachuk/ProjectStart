package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /tracks
// Получаем список всех треков
func GetAllTracks(context *gin.Context) {
	var tracks []models.Track
	models.DB.Find(&tracks)

	context.JSON(http.StatusOK, gin.H{"tracks": tracks})
}
