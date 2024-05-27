package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "web-service-gin/models"
)

func GetAlbums(c *gin.Context) { // Correct capitalization
    var albums []models.Album
    if err := models.DB.Find(&albums).Error; err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve albums"})
        return
    }
    c.IndentedJSON(http.StatusOK, albums)
}
