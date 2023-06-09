package adminservice

import (
	webadmin "pancakaki/internal/domain/web/admin"
	webbank "pancakaki/internal/domain/web/bank"
)

type AdminService interface {
	Register(req webadmin.AdminCreateRequest) (webadmin.AdminResponse, error)
	ViewAll() ([]webadmin.AdminResponse, error)
	ViewOne(adminId int, username string) (webadmin.AdminResponse, error)
	Edit(req webadmin.AdminUpdateRequest) (webadmin.AdminResponse, error)
	Unreg(adminId int, username string) (webadmin.AdminResponse, error)

	RegisterBank(req webbank.BankCreateRequest, reqBank webbank.BankAdminCreateRequest) (webbank.BankResponse, error)
	ViewAllBank() ([]webbank.BankResponse, error)
	EditBank(req webbank.BankUpdateRequest) (webbank.BankResponse, error)
	DeleteBank(bankId int) ([]webbank.BankResponse, error)

	// ViewTransactionAllOwner() ([]webadmin.TransactionOwnerResponse, error)
	// ViewTransactionOwnerByName(ownerName string) (webadmin.TransactionOwnerResponse, error)

	// ViewAllOwner() ([]webadmin.FindOwnerResponse, error)
	// ViewOwnerByName(ownerName string) (webadmin.FindOwnerResponse, error)
	// UnregOwner(ownerId int) (webowner.OwnerResponse, error)

	// ViewTransactionCustomerById(customerId int) ([]webcustomer.TransactionCustomer, error)
}
