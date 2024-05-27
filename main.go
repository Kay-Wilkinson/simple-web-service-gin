package main

import (
    "github.com/gin-gonic/gin"

    "web-service-gin/handlers"
    "web-service-gin/models"
)

func main() {
    models.ConnectDatabase() // Correct capitalization

    router := gin.Default()
    router.GET("/albums", handlers.GetAlbums)       // Correct capitalization
    router.GET("/albums/:id", handlers.GetAlbumByID) // Correct capitalization
    router.POST("/albums", handlers.PostAlbums)      // Correct capitalization

    router.Run("0.0.0.0:8080")
}
