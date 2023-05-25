package webtransaction

import "time"

type TransactionResponse struct {
	ProductName string    `json:"product_nmame"`
	Qty         int       `json:"quantity"`
	Tax         float64   `json:"tax"`
	TotalPrice  int64     `json:"total_price"`
	BuyDate     time.Time `json:"buy_date"`
}
