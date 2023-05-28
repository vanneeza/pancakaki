package webstore

type StoreUpdateRequest struct {
	Id          int    `json:"store_id"`
	Name        string `json:"name"`
	NoHp        string `json:"no_hp"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	OwnerId     int    `json:"owner_id"`
	BankId      int    `json:"bank_id"`
	BankName    string `json:"bankName"`
	BankAccount int64  `json:"bankAccount"`
	AccountName string `json:"accountName"`
}
