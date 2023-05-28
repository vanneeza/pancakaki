package webowner

import "pancakaki/internal/domain/entity"

type OwnerUpdateResponse struct {
	Id              int           `json:"id"`
	Name            string        `json:"name"`
	NoHp            string        `json:"no_hp"`
	Email           string        `json:"email"`
	Password        string        `json:"passowrd"`
	MembershipName  string        `json:"membershipName"`
	MembershipPrice int64         `json:"membershipPrice"`
	Bank            []entity.Bank `json:"bank"`
}
