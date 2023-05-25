package webowner

type OwnerResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	NoHp         string `json:"no_hp"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MembershipId int    `json:"membership_id"`
}
