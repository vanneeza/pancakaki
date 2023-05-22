package entity

import "time"

type Product struct {
	Id          int
	Name        string
	Price       int
	Stock       int16
	Description string
	CreatedAt   time.Time
	UpdateAt    time.Time
	DiscountId  int
	MerkId      int
}
