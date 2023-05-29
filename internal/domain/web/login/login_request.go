package weblogin

type LoginRequest struct {
	Username string `json:"username"`
	NoHp     string `json:"no_hp"`
	Password string `json:"password"`
}
