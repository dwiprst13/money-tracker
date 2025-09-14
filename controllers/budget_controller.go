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

func UpdateBudget(c *gin.Context) {
	userID := c.GetString("userID")
	budgetID := c.Param("id")

	var budget models.Budget
	if err := config.DB.Where("id = ? AND user_id = ?", budgetID, userID).First(&budget).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
		return
	}

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

	budget.Category = input.Category
	budget.Amount = input.Amount
	budget.StartDate = parseDate(input.StartDate)
	budget.EndDate = parseDate(input.EndDate)

	config.DB.Save(&budget)
	c.JSON(http.StatusOK, budget)
}

func DeleteBudget(c *gin.Context) {
	userID := c.GetString("userID")
	budgetID := c.Param("id")

	var budget models.Budget
	if err := config.DB.Where("id = ? AND user_id = ?", budgetID, userID).First(&budget).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
		return
	}

	config.DB.Delete(&budget)
	c.JSON(http.StatusOK, gin.H{"message": "Budget deleted successfully"})
}

func parseDate(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{} 
	}
	return t
}

