package customerservice

import (
	webcustomer "pancakaki/internal/domain/web/customer"
	webtransaction "pancakaki/internal/domain/web/transaction"
)

type CustomerService interface {
	Register(req webcustomer.CustomerCreateRequest) (webcustomer.CustomerResponse, error)
	ViewAll() ([]webcustomer.CustomerResponse, error)
	ViewOne(customerId int, customerName, customerNoHp string) (webcustomer.CustomerResponse, error)
	Edit(req webcustomer.CustomerUpdateRequest) (webcustomer.CustomerResponse, error)
	Unreg(customerId int, customerName, customerNoHp string) (webcustomer.CustomerResponse, error)
	Notification(customerId int) ([]webtransaction.TransactionResponse, error)
}
