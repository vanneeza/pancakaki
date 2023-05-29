package entity

import "time"

type TransactionOrderDetail struct {
	Id             int
	BuyDate        time.Time
	Status         string
	TotalPrice     int
	Photo          string
	Tax            float64
	VirtualAccount int64
}

type TransactionOrder struct {
	Id            int
	Qty           int
	Total         int
	CustomerId    int
	ProductId     int
	DetailOrderId int
}

type Payment struct {
	Id                       int
	TransactionDetailOrderId int
	Pay                      float64
}
