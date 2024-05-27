package tests

import (
	"fmt"
	"net/http"
    "net/http/httptest"
    "testing"
    "web-service-gin/handlers"
    "web-service-gin/models"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "encoding/json"
)

// SetupTestDatabase: Initializes a test database connection and applies migrations.
// TestGetAlbumByID:
// Initializes the test database and assigns it to the global models.DB.
// Inserts a test album directly into the database.
// Sets up the Gin router and the handler.
// Sends a request to the API endpoint.
// Queries the same album directly from the database.
// Compares the API response with the expected JSON string.


func TestGetAlbumByID(t *testing.T) {
    // Initialize the test database
    db, err := SetupTestDatabase()
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }
    defer TeardownTestDatabase(db)
    models.DB = db

    // Create a test album
    testAlbum := models.Album{Title: "Test Album", Artist: "Test Artist", Price: 19.99}
    db.Create(&testAlbum)

    // Set up the router and the handler
    router := gin.Default()
    router.GET("/albums/:id", handlers.GetAlbumByID)

    req, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%d", testAlbum.ID), nil)

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    var albumFromDB models.Album
    db.First(&albumFromDB, "id = ?", testAlbum.ID)

    var responseAlbum models.Album
    json.Unmarshal(w.Body.Bytes(), &responseAlbum)

    assert.Equal(t, albumFromDB, responseAlbum)
}






