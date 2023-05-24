package webmembership

type MembershipResponse struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Tax   float64 `json:"tax"`
	Price int64   `json:"price"`
}
