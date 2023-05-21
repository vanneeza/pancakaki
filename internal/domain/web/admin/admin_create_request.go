package webadmin

type AdminCreateRequest struct {
	Name     string `json:"name"`
	Passowrd string `json:"password"`
}
