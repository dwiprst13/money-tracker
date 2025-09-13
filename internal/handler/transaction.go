package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"money-tracker/models"
)

type TransactionHandler struct {
	DB *gorm.DB
}

func (h *TransactionHandler) List(c *gin.Context) {
	var transactions []models.Transaction
	h.DB.Order("date(created_at) desc").Order("time(created_at) desc").Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) Create(c *gin.Context) {
	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.CreatedAt = time.Now()
	if err := h.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func (h *TransactionHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.DB.Delete(&models.Transaction{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
