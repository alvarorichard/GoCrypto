package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Coin is the coin model
type Coin struct {
	Id        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name" binding:"required"`
	OwnerID   string    `gorm:"index"`
	Symbol    string    `json:"symbol" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
}

func (c *Coin) BeforeCreate(tx *gorm.DB) (err error) {
	c.Id = uuid.New().String()
	c.CreatedAt = time.Now()
	return
}
