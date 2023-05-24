package webmembership

type MembershipCreateRequest struct {
	Name  string  `json:"name"`
	Tax   float64 `json:"tax"`
	Price int64   `json:"price"`
}
