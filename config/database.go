package config

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "money-tracker/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	
    dsn := "root:password@tcp(127.0.0.1:3306)/money_tracker?charset=utf8mb4&parseTime=True&loc=Local"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to MySQL database: %v", err))
    }

    database.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Budget{})

    DB = database
}
