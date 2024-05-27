package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "web-service-gin/models"
)

func PostAlbums(c *gin.Context) {
    var newAlbum models.Album

    if err := c.BindJSON(&newAlbum); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }

    if err := models.DB.Create(&newAlbum).Error; err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create album"})
        return
    }
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
