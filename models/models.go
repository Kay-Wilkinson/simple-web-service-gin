package models

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "fmt"
)

var DB *gorm.DB

type Album struct {
    ID     int     `json:"id" gorm:"primaryKey;autoIncrement"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

func ConnectDatabase() { // Correct capitalization
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    db.AutoMigrate(&Album{})

    DB = db
}
