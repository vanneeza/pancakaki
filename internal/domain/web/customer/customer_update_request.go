package webcustomer

type CustomerUpdateRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NoHp     string `json:"no_hp"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
