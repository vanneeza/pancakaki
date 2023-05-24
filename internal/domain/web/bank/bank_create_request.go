package webbank

type BankCreateRequest struct {
	Name        string `json:"name"`
	BankAccount int64  `json:"bank_account"`
	AccountName string `json:"account_name"`
}

type BankAdminCreateRequest struct {
	AdminId int `json:"admin_id"`
}
