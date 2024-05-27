package tests

import (
    "web-service-gin/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)


func SetupTestDatabase() (*gorm.DB, error) {
    dsn := "host=db user=dev_user password=dev_password dbname=dev_dbname port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }


    db.AutoMigrate(&models.Album{})
    return db, nil
}


func TeardownTestDatabase(db *gorm.DB) error {
    return db.Migrator().DropTable(&models.Album{})
}
