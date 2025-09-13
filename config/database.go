package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "money-tracker/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Budget{})
    DB = database
}
