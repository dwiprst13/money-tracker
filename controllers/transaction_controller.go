package controllers

import (
	"net/http"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"money-tracker/config"
	"money-tracker/models"
)

func GetTransactions(c *gin.Context) {
	// userID := c.GetString("userID")
	var transactions []models.Transaction
	config.DB.Where("user_id = ?", 1).Order("date desc").Find(&transactions)
	// config.DB.Where("user_id = ?", userID).Order("date desc").Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	// userID := c.GetString("userID")

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
		// UserID:      parseUint(userID),
		UserID:      1,
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
