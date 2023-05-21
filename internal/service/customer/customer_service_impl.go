package customerservice

import (
	"fmt"
	"log"
	entity "pancakaki/internal/domain/entity/customer"
	webcustomer "pancakaki/internal/domain/web/customer"
	customerrepository "pancakaki/internal/repository/customer"
	"pancakaki/utils/helper"
)

type CustomerServiceImpl struct {
	CustomerRepository customerrepository.CustomerRepository
}

func NewCustomerService(customerRepository customerrepository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
	}
}

func (customerService *CustomerServiceImpl) Register(req webcustomer.CustomerCreateRequest) (webcustomer.CustomerResponse, error) {

	customer := entity.Customer{
		Name:    req.Name,
		NoHp:    req.NoHp,
		Address: req.Address,
		Photo:   req.Photo.Filename,
		Balance: req.Balance,
	}

	log.Println(customer, "Ini data Customer")
	fmt.Scanln()
	customerData, _ := customerService.CustomerRepository.Create(&customer)

	log.Println(customerData, "Ini data Customer Data")
	fmt.Scanln()
	customerResponse := webcustomer.CustomerResponse{
		Id:      customerData.Id,
		Name:    customerData.Name,
		NoHp:    customerData.NoHp,
		Address: customerData.Address,
		Photo:   customerData.Photo,
		Balance: customerData.Balance,
	}
	return customerResponse, nil
}

func (customerService *CustomerServiceImpl) ViewAll() ([]webcustomer.CustomerResponse, error) {
	customerData, err := customerService.CustomerRepository.FindAll()
	helper.PanicErr(err)

	customerResponse := make([]webcustomer.CustomerResponse, len(customerData))
	for i, customer := range customerData {
		customerResponse[i] = webcustomer.CustomerResponse{
			Id:      customer.Id,
			Name:    customer.Name,
			NoHp:    customer.NoHp,
			Address: customer.Address,
			Photo:   customer.Photo,
			Balance: customer.Balance,
		}
	}
	return customerResponse, nil
}

func (customerService *CustomerServiceImpl) ViewOne(customerId int) (webcustomer.CustomerResponse, error) {
	customer, err := customerService.CustomerRepository.FindById(customerId)
	helper.PanicErr(err)

	customerResponse := webcustomer.CustomerResponse{
		Id:      customerId,
		Name:    customer.Name,
		NoHp:    customer.NoHp,
		Address: customer.Address,
		Photo:   customer.Photo,
		Balance: customer.Balance,
	}

	return customerResponse, nil
}

func (customerService *CustomerServiceImpl) Edit(req webcustomer.CustomerUpdateRequest) (webcustomer.CustomerResponse, error) {

	customer := entity.Customer{
		Id:      req.Id,
		Name:    req.Name,
		NoHp:    req.NoHp,
		Address: req.Address,
		Photo:   req.Photo,
		Balance: req.Balance,
	}

	customerData, err := customerService.CustomerRepository.Update(&customer)
	helper.PanicErr(err)

	customerResponse := webcustomer.CustomerResponse{
		Id:      customerData.Id,
		Name:    customerData.Name,
		NoHp:    customerData.NoHp,
		Address: customerData.Address,
		Photo:   customerData.Photo,
		Balance: customerData.Balance,
	}

	return customerResponse, nil
}

func (customerService *CustomerServiceImpl) Unreg(customerId int) (webcustomer.CustomerResponse, error) {

	customerData, err := customerService.CustomerRepository.FindById(customerId)
	helper.PanicErr(err)

	err = customerService.CustomerRepository.Delete(customerId)
	helper.PanicErr(err)

	customerResponse := webcustomer.CustomerResponse{
		Name:    customerData.Name,
		NoHp:    customerData.NoHp,
		Address: customerData.Address,
		Photo:   customerData.Photo,
		Balance: customerData.Balance,
	}

	return customerResponse, nil
}
