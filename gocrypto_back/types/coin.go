package types

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	Name    string    `json:"name"`
	Balance []Balance `json:"balance" gorm:"foreignKey:CoinID"`
	Symbol  string    `json:"symbol"`
	Price   int       `json:"price"`
}

type NewCoin struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Price  int    `json:"price"`
}
