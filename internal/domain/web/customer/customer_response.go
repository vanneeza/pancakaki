package webcustomer

type CustomerResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	NoHp    int64  `json:"no_hp"`
	Address string `json:"address"`
	Photo   string `json:"photo"`
	Balance int64  `json:"balance"`
}
