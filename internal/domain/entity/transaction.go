package entity

import "time"

type TransactionOrderDetail struct {
	Id         int
	BuyDate    time.Time
	Status     string
	TotalPrice int64
	Photo      string
	Tax        float64
}

type TransactionOrder struct {
	Id            int
	Qty           int
	Total         int
	CustomerId    int
	ProductId     int
	DetailOrderId int
}
