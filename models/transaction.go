package models

import "time"

type Transaction struct {
    ID         uint      `gorm:"primaryKey"`
    UserID     uint      `gorm:"index"`
    Amount     float64
    Category   string
    Note       string
    ReceiptURL string
    CreatedAt  time.Time
}
