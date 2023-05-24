package webbank

type BankResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	AccountName string `json:"account_name"`
	BankAccount int64  `json:"bank_account"`
}
