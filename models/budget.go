package models

type Budget struct {
    ID       uint    `gorm:"primaryKey"`
    UserID   uint    `gorm:"index"`
    Category string
    Amount   float64
}
