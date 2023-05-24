package entity

import "time"

type Admin struct {
	Id       int
	Username string
	Password string
}

type TransactionOwner struct {
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

type FindOwner struct {
	OwnerName      string
	NoHp           int64
	Email          string
	Password       string
	NameMembership string
	NameStore      string
}
