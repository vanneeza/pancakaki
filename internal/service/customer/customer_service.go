package customerservice

import (
	webcustomer "pancakaki/internal/domain/web/customer"
	webtransaction "pancakaki/internal/domain/web/transaction"
)

type CustomerService interface {
	Register(req webcustomer.CustomerCreateRequest) (webcustomer.CustomerResponse, error)
	ViewAll() ([]webcustomer.CustomerResponse, error)
	ViewOne(customerName string) (webcustomer.CustomerResponse, error)
	ViewByNoHp(noHp string) (webcustomer.CustomerResponse, error)
	Edit(req webcustomer.CustomerUpdateRequest) (webcustomer.CustomerResponse, error)
	Unreg(customerName string) (webcustomer.CustomerResponse, error)
	Notification(customerId int) ([]webtransaction.TransactionResponse, error)
}
