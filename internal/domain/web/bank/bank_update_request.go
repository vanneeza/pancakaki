package webbank

type BankUpdateRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	BankAccount int64  `json:"bank_account"`
	AccountName string `json:"account_name"`
}

type BankAdminUpdateRequest struct {
	AdminId int `json:"admin_id"`
}
