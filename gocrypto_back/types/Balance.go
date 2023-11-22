package types

import (
	"time"

	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	UserID    uint     `json:"user_id"`
	CoinID    uint     `json:"coin_id"`
	Amount    float64  `json:"amount"`
	Operation string   `json:"op"`
	Balance   *Balance `gorm:"foreignkey:UserID"` // Assegure que a relação está definida corretamente

}

type WalletResponse struct {
	Date      time.Time `json:"time"`
	Ammount   float64   `json:"amount"`
	CoinName  string    `json:"coin"`
	Operation string    `json:"operation"`
}
