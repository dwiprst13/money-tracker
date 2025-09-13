package controllers

import (
    "net/http"
    "time"
    "money-tracker/config"
    "money-tracker/models"
    "money-tracker/utils"

    "github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
    userID := c.GetString("userID")
    var input struct {
        Amount     float64 `json:"amount" binding:"required"`
        Category   string  `json:"category" binding:"required"`
        Note       string  `json:"note"`
        ReceiptURL string  `json:"receipt_url"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    tx := models.Transaction{
        UserID:     utils.ParseUint(userID),
        Amount:     input.Amount,
        Category:   input.Category,
        Note:       input.Note,
        ReceiptURL: input.ReceiptURL,
        CreatedAt:  time.Now(),
    }
    if err := config.DB.Create(&tx).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan transaksi"})
        return
    }
    c.JSON(http.StatusOK, tx)
}

func ListTransactions(c *gin.Context) {
    userID := c.GetString("userID")
    var transactions []models.Transaction
    config.DB.Where("user_id = ?", utils.ParseUint(userID)).Order("created_at desc").Find(&transactions)
    c.JSON(http.StatusOK, transactions)
}
