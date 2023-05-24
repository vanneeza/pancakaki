package webowner

import entity "pancakaki/internal/domain/entity/bank"

type OwnerCreateResponse struct {
	Id              int
	Name            string
	NoHp            string
	Email           string
	Password        string
	MembershipName  string
	MembershipPrice int
	Bank            []entity.Bank
	Token           string
	// BankAccount     int64
	// AccountName     string
}
