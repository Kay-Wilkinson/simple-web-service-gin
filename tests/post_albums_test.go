package tests

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "web-service-gin/handlers"
    "web-service-gin/models"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)


func TestPostAlbums(t *testing.T) {
    db, err := SetupTestDatabase()
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }
    defer TeardownTestDatabase(db)
    models.DB = db

    router := gin.Default()
    router.POST("/albums", handlers.PostAlbums)

    newAlbum := []byte(`{"title":"Test Album","artist":"Test Artist","price":77.99}`)
    req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(newAlbum))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusCreated {
        t.Logf("Response body: %s", w.Body.String())
    }
    assert.Equal(t, http.StatusCreated, w.Code)


    assert.Contains(t, w.Body.String(), "Test Album")
}
