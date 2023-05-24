package webadmin

type AdminCreateRequest struct {
	Username string `json:"username"`
	Passowrd string `json:"password"`
}
