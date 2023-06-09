package webcustomer

type CustomerResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NoHp     string `json:"no_hp"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type CustomerRegisterResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NoHp     int64  `json:"no_hp"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
