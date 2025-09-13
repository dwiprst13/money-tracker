package controllers

import (
    "net/http"
    "money-tracker/config"
    "money-tracker/models"

    "github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := models.User{Email: input.Email, Password: input.Password}
    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah terdaftar"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Register berhasil"})
}

func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := config.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "userID": user.ID})
}
