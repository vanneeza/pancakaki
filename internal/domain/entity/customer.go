package entity

import "time"

type Customer struct {
	Id       int
	Name     string
	NoHp     string
	Address  string
	Password string
	IsDelete bool
	Role     string
}

type TransactionCustomer struct {
	CustomerName   string
	MerkName       string
	ProductId      int
	ProductName    string
	ProductPrice   int
	ShippingCost   int
	Qty            int
	Tax            float64
	TotalPrice     float64
	BuyDate        time.Time
	Status         string
	StoreName      string
	VirtualAccount int
}
