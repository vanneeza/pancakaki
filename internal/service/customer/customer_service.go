package customerservice

import (
	webcustomer "pancakaki/internal/domain/web/customer"
)

type CustomerService interface {
	Register(req webcustomer.CustomerCreateRequest) (webcustomer.CustomerResponse, error)
	ViewAll() ([]webcustomer.CustomerResponse, error)
	ViewOne(customerName string) (webcustomer.CustomerResponse, error)
	Edit(req webcustomer.CustomerUpdateRequest) (webcustomer.CustomerResponse, error)
	Unreg(customerName string) (webcustomer.CustomerResponse, error)
}
