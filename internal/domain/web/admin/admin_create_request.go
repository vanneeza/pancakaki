package webadmin

type AdminCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
