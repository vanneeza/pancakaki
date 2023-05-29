package webowner

type OwnerCreateResponse struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	NoHp            string `json:"no_hp"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	MembershipName  string `json:"membershipName"`
	MembershipPrice int64  `json:"membershipPrice"`
	VirtualAccount  int    `json:"virtual_account"`
	Token           string `json:"token"`
}
