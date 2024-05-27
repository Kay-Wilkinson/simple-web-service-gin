package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "web-service-gin/models"
)

func GetAlbumByID(c *gin.Context) { // Correct capitalization
    id := c.Param("id")
    var album models.Album

    if err := models.DB.First(&album, "id = ?", id).Error; err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
        return
    }
    c.IndentedJSON(http.StatusOK, album)
}
