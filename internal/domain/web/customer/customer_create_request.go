package webcustomer

type CustomerCreateRequest struct {
	Name     string `json:"name"`
	NoHp     string `json:"no_hp"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
