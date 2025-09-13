package controllers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "money-tracker/config"
    "money-tracker/models"
    "money-tracker/utils"
)

func GetReports(c *gin.Context) {
    // userID := c.GetString("userID")
    userID := "1"

    var totalIncome float64
    var totalExpense float64

    startOfMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)
    endOfMonth := startOfMonth.AddDate(0, 1, 0)

    var transactions []models.Transaction
    config.DB.Where("user_id = ? AND created_at >= ? AND created_at < ?", utils.ParseUint(userID), startOfMonth, endOfMonth).Find(&transactions)

    for _, t := range transactions {
        if t.Amount > 0 {
            totalIncome += t.Amount
        } else {
            totalExpense += t.Amount
        }
    }

    saldo := totalIncome + totalExpense

    insight := "Bulan ini pengeluaran makan naik 20% dibanding bulan lalu."

    c.JSON(http.StatusOK, gin.H{
        "total_income":  totalIncome,
        "total_expense": totalExpense,
        "saldo":         saldo,
        "insight":       insight,
    })
}
