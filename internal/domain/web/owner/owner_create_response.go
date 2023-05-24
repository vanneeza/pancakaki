package webowner

import "pancakaki/internal/domain/entity"

type OwnerCreateResponse struct {
	Id              int
	Name            string
	NoHp            string
	Email           string
	Password        string
	MembershipName  string
	MembershipPrice int64
	Bank            []entity.Bank
	Token           string
	// BankAccount     int64
	// AccountName     string
}
