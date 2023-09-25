package models

import "gorm.io/gorm"

// Coin is the coin model
type Coin struct {
	gorm.Model
	Name   string
	Symbol string
	Price  float64
}
