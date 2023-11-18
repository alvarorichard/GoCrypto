package types

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	UserID uint    `json:"user_id"`
	CoinID uint    `json:"coin_id"`
	Amount float64 `json:"amount"`
}
