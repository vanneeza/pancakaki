package entity

import "time"

type Customer struct {
	Id       int
	Name     string
	NoHp     int64
	Address  string
	Password string
	IsDelete bool
}

type TransactionCustomer struct {
	CustomerName   string
	MerkName       string
	ProductName    string
	ProductPrice   int
	ShippingCost   int
	Qty            int
	Tax            float64
	TotalPrice     int
	BuyDate        time.Time
	Status         string
	StoreName      string
	VirtualAccount int
}
