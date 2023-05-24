package webstore

type StoreCreateRequest struct {
	Name        string
	NoHp        int
	Email       string
	Address     string
	OwnerId     int
	BankName    string
	BankAccount int64
	AccountName string
}
