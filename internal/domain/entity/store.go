package entity

import "time"

type Store struct {
	Id        int
	Name      string
	NoHp      int
	Email     string
	Address   string
	OwnerId   int
	IsDeleted bool
}

type TransactionStore struct {
	Id             int
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
