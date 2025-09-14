package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"money-tracker/config"
	"money-tracker/models"
)

func GetTransactions(c *gin.Context) {
	userID := c.GetString("userID")
	var transactions []models.Transaction
	config.DB.Where("user_id = ?", userID).Order("date desc").Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	userID := c.GetString("userID")

	var input struct {
		Type        string  `json:"type"`
		Category    string  `json:"category"`
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := models.Transaction{
		UserID:      parseUint(userID),
		Type:        input.Type,
		Category:    input.Category,
		Amount:      input.Amount,
		Description: input.Description,
		Date:        time.Now(),
	}

	config.DB.Create(&tx)
	c.JSON(http.StatusOK, tx)
}

func parseUint(s string) uint {
	var id uint
	fmt.Sscanf(s, "%d", &id)
	return id
}

func UpdateTransaction(c *gin.Context) {
	userID := c.GetString("userID")
	txID := c.Param("id")

	var tx models.Transaction
	if err := config.DB.Where("id = ? AND user_id = ?", txID, userID).First(&tx).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	var input struct {
		Type        string  `json:"type"`
		Category    string  `json:"category"`
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx.Type = input.Type
	tx.Category = input.Category
	tx.Amount = input.Amount
	tx.Description = input.Description

	config.DB.Save(&tx)
	c.JSON(http.StatusOK, tx)
}

func DeleteTransaction(c *gin.Context) {
	userID := c.GetString("userID")
	txID := c.Param("id")

	var tx models.Transaction
	if err := config.DB.Where("id = ? AND user_id = ?", txID, userID).First(&tx).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	config.DB.Delete(&tx)
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
