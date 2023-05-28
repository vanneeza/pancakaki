package webcustomer

import "time"

type CustomerResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NoHp     int64  `json:"no_hp"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type TransactionCustomer struct {
	CustomerName string    `json:"customer_name"`
	ProductName  string    `json:"product_name"`
	MerkName     string    `json:"merk_name"`
	Price        float64   `json:"price"`
	Qty          int       `json:"quantity"`
	BuyDate      time.Time `json:"buy_date"`
	TotalPrice   int64     `json:"total_price"`
	Status       string    `json:"status"`
	OwnerName    string    `json:"owner_name"`
}
