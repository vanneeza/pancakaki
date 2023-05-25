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
	NameProduct  string
	NameMerk     string
	Price        float64
	Qty          int
	BuyDate      time.Time
	TotalPrice   int64
	Status       string
	CustomerName string
	OwnerName    string
}
