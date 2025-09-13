package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"money-tracker/config"
	"money-tracker/models"
)

func GetBudgets(c *gin.Context) {
	userID := c.GetString("userID")
	var budgets []models.Budget
	config.DB.Where("user_id = ?", userID).Find(&budgets)
	c.JSON(http.StatusOK, budgets)
}

func CreateBudget(c *gin.Context) {
	userID := c.GetString("userID")

	var input struct {
		Category  string  `json:"category"`
		Amount    float64 `json:"amount"`
		StartDate string  `json:"start_date"`
		EndDate   string  `json:"end_date"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	budget := models.Budget{
		UserID:    parseUint(userID),
		Category:  input.Category,
		Amount:    input.Amount,
		StartDate: parseDate(input.StartDate),
		EndDate:   parseDate(input.EndDate),
	}

	config.DB.Create(&budget)
	c.JSON(http.StatusOK, budget)
}

func parseDate(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s)
	return t
}
