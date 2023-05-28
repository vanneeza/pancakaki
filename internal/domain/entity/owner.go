package entity

import "time"

type Owner struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	NoHp         string `json:"no_hp"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MembershipId int    `json:"membership_id"`
	Role         string `json:"role"`
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
