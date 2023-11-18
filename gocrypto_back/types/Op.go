package types

type Operation struct {
	CoinID    uint    `json:"coin_id"`
	Amount    float64 `json:"amount"`
	Operation string  `json:"operation"`
}
