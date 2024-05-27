package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"web-service-gin/handlers"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestGetAlbums(t *testing.T) {
	db, err := SetupTestDatabase()
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }
    models.DB = db

    // Create a test album
    testAlbum := models.Album{Title: "Blue Train", Artist: "Blue Artist", Price: 9.99}
    db.Create(&testAlbum)

	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)

	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Blue Train")
}




