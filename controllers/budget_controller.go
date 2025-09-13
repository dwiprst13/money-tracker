package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "money-tracker/config"
    "money-tracker/models"
    "money-tracker/utils"
)

func SetBudget(c *gin.Context) {
    userID := c.GetString("userID")

    var input struct {
        Category string  `json:"category" binding:"required"`
        Amount   float64 `json:"amount" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var budget models.Budget
    err := config.DB.Where("user_id = ? AND category = ?", utils.ParseUint(userID), input.Category).First(&budget).Error
    if err != nil {
        budget = models.Budget{
            UserID:   utils.ParseUint(userID),
            Category: input.Category,
            Amount:   input.Amount,
        }
        config.DB.Create(&budget)
    } else {
        budget.Amount = input.Amount
        config.DB.Save(&budget)
    }

    c.JSON(http.StatusOK, budget)
}

func GetBudgets(c *gin.Context) {
    userID := c.GetString("userID")
    var budgets []models.Budget
    config.DB.Where("user_id = ?", utils.ParseUint(userID)).Find(&budgets)
    c.JSON(http.StatusOK, budgets)
}
