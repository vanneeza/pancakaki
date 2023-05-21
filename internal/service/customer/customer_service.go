package customerservice

import (
	webcustomer "pancakaki/internal/domain/web/customer"
)

type CustomerService interface {
	Register(req webcustomer.CustomerCreateRequest) (webcustomer.CustomerResponse, error)
	ViewAll() ([]webcustomer.CustomerResponse, error)
	ViewOne(customerId int) (webcustomer.CustomerResponse, error)
	Edit(req webcustomer.CustomerUpdateRequest) (webcustomer.CustomerResponse, error)
	Unreg(customerId int) (webcustomer.CustomerResponse, error)
}
