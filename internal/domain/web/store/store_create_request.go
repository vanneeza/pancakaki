package webstore

type StoreCreateRequest struct {
	Name        string `json:"name"`
	NoHp        string `json:"no_hp"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	OwnerId     int    `json:"owner_id"`
	BankName    string `json:"bank_name"`
	BankAccount int64  `json:"bank_account"`
	AccountName string `json:"account_name"`
}
