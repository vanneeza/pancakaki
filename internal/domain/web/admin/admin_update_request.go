package webadmin

type AdminUpdateRequest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
