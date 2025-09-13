package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"money-tracker/models"
)

type ReportHandler struct {
	DB *gorm.DB
}

func (h *ReportHandler) Monthly(c *gin.Context) {
	var totalIncome float64
	var totalExpense float64

	startOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	var transactions []models.Transaction
	h.DB.Where("created_at >= ? AND created_at < ?", startOfMonth, endOfMonth).Find(&transactions)

	for _, t := range transactions {
		if t.Amount > 0 {
			totalIncome += t.Amount
		} else {
			totalExpense += t.Amount
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"total_income":  totalIncome,
		"total_expense": totalExpense,
		"saldo":         totalIncome + totalExpense,
	})
}
