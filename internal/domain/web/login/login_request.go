package weblogin

type LoginRequest struct {
	NoHp     string `json:"no_hp"`
	Password string `json:"password"`
}
