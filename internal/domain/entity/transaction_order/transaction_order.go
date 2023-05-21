package entity

import "time"

type TransactionOrder struct {
	Id         int
	Qty        int
	BuyDate    time.Time
	Status     string
	Total      int64
	CustomerId int
	ProductId  int
	PacketId   int
}
