package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"money-tracker/models"
)

type BudgetHandler struct {
	DB *gorm.DB
}

func (h *BudgetHandler) List(c *gin.Context) {
	var budgets []models.Budget
	h.DB.Find(&budgets)
	c.JSON(http.StatusOK, budgets)
}

func (h *BudgetHandler) Set(c *gin.Context) {
	var input models.Budget
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var budget models.Budget
	err := h.DB.Where("user_id = ? AND category = ?", input.UserID, input.Category).First(&budget).Error
	if err != nil {
		h.DB.Create(&input)
		c.JSON(http.StatusCreated, input)
		return
	}
	budget.Amount = input.Amount
	h.DB.Save(&budget)
	c.JSON(http.StatusOK, budget)
}
