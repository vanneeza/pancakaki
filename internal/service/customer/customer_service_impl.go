package customerservice

import (
	"fmt"
	"log"
	"pancakaki/internal/domain/entity"
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
		Name:     req.Name,
		NoHp:     req.NoHp,
		Address:  req.Address,
		Password: req.Password,
	}
	log.Println(customer, "service reg")
	log.Println(req, "service reg")
	fmt.Scanln()
	customerData, _ := customerService.CustomerRepository.Create(&customer)

	customerResponse := webcustomer.CustomerResponse{
		Id:       customerData.Id,
		Name:     customerData.Name,
		NoHp:     customerData.NoHp,
		Address:  customerData.Address,
		Password: customerData.Password,
	}

	return customerResponse, nil
}

func (customerService *CustomerServiceImpl) ViewAll() ([]webcustomer.CustomerResponse, error) {
	customerData, err := customerService.CustomerRepository.FindAll()
	helper.PanicErr(err)

	customerResponse := make([]webcustomer.CustomerResponse, len(customerData))
	for i, customer := range customerData {
		customerResponse[i] = webcustomer.CustomerResponse{
			Id:       customer.Id,
			Name:     customer.Name,
			NoHp:     customer.NoHp,
			Address:  customer.Address,
			Password: customer.Password,
		}
	}
	return customerResponse, nil
}

func (customerService *CustomerServiceImpl) ViewOne(customerName string) (webcustomer.CustomerResponse, error) {
	customer, err := customerService.CustomerRepository.FindByName(customerName)
	helper.PanicErr(err)

	customerResponse := webcustomer.CustomerResponse{
		Id:       customer.Id,
		Name:     customer.Name,
		NoHp:     customer.NoHp,
		Address:  customer.Address,
		Password: customer.Password,
	}

	return customerResponse, nil
}

func (customerService *CustomerServiceImpl) Edit(req webcustomer.CustomerUpdateRequest) (webcustomer.CustomerResponse, error) {

	customer := entity.Customer{
		Id:       req.Id,
		Name:     req.Name,
		NoHp:     req.NoHp,
		Address:  req.Address,
		Password: req.Password,
	}

	customerData, err := customerService.CustomerRepository.Update(&customer)
	helper.PanicErr(err)

	customerResponse := webcustomer.CustomerResponse{
		Id:       customerData.Id,
		Name:     customerData.Name,
		NoHp:     customerData.NoHp,
		Address:  customerData.Address,
		Password: customerData.Password,
	}

	return customerResponse, nil
}

func (customerService *CustomerServiceImpl) Unreg(customerName string) (webcustomer.CustomerResponse, error) {

	customerData, err := customerService.CustomerRepository.FindByName(customerName)
	helper.PanicErr(err)

	err = customerService.CustomerRepository.Delete(customerData.Id)
	helper.PanicErr(err)

	customerResponse := webcustomer.CustomerResponse{
		Name:     customerData.Name,
		NoHp:     customerData.NoHp,
		Address:  customerData.Address,
		Password: customerData.Password,
	}

	return customerResponse, nil
}
