package entity

type Bank struct {
	Id          int
	Name        string
	BankAccount int64
	AccountName string
}

type BankAdmin struct {
	Id      int
	AdminId int
	BankId  int
}

type BankStore struct {
	Id      int
	StoreId int
	BankId  int
}
