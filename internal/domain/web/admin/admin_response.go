package webadmin

type AdminResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Passowrd string `json:"password"`
}
