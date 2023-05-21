package webadmin

type AdminUpdateRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Passowrd string `json:"password"`
}
