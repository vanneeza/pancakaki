package webtransaction

import "time"

type TransactionOrderDetailCreateRequest struct {
	BuyDate    time.Time `json:"buy_date"`
	Status     string    `json:"status"`
	TotalPrice int       `json:"total_price"`
	Photo      string    `json:"photo"`
	Tax        float64   `json:"tax"`
}

type TransactionOrderCreateRequest struct {
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"product_id"`
	Qty        int `json:"qty"`
}
