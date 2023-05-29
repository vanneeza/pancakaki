package webadmin

import "time"

type AdminResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type TransactionOwnerResponse struct {
	OwnerName    string    `json:"owner_name"`
	ProductName  string    `json:"product_name"`
	MerkName     string    `json:"merk_name"`
	Price        float64   `json:"price"`
	Qty          int       `json:"quantity"`
	BuyDate      time.Time `json:"buy_date"`
	TotalPrice   int64     `json:"total_price"`
	Status       string    `json:"status"`
	CustomerName string    `json:"customer_name"`
}

type FindOwnerResponse struct {
	OwnerName      string `json:"owner_name"`
	NoHp           int64  `json:"no_hp"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	NameMembership string `json:"membership_name"`
	NameStore      string `json:"store_name"`
}
