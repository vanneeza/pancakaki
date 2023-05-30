package webowner

import (
	webmembership "pancakaki/internal/domain/web/membership"
)

type OwnerResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	NoHp         string `json:"no_hp"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MembershipId int    `json:"membership_id"`
}

type GetOwnerResponse struct {
	Id         int                              `json:"id"`
	Name       string                           `json:"name"`
	NoHp       string                           `json:"no_hp"`
	Email      string                           `json:"email"`
	Password   string                           `json:"password"`
	Membership webmembership.MembershipResponse `json:"membership"`
}
