package webowner

import "pancakaki/internal/domain/entity"

type OwnerCreateResponse struct {
	Id              int           `json:"id"`
	Name            string        `json:"name"`
	NoHp            string        `json:"no_hp"`
	Email           string        `json:"email"`
	Password        string        `json:"passowrd"`
	MembershipName  string        `json:"membershipName"`
	MembershipPrice int64         `json:"membershipPrice"`
	Bank            []entity.Bank `json:"bank"`
	Token           string        `json:"token"`
	// BankAccount     int64
	// AccountName     string
}
