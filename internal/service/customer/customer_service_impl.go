package customerservice

import (
	"fmt"
	"pancakaki/internal/domain/entity"
	webcustomer "pancakaki/internal/domain/web/customer"
	webtransaction "pancakaki/internal/domain/web/transaction"
	customerrepository "pancakaki/internal/repository/customer"
	"pancakaki/utils/helper"

	"golang.org/x/crypto/bcrypt"
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

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	customer := entity.Customer{
		Name:     req.Name,
		NoHp:     req.NoHp,
		Address:  req.Address,
		Password: string(encryptedPassword),
	}

	customerData, _ := customerService.CustomerRepository.Create(&customer)

	customerResponse := webcustomer.CustomerResponse{
		Id:       customerData.Id,
		Name:     customerData.Name,
		NoHp:     customerData.NoHp,
		Address:  customerData.Address,
		Password: customerData.Password,
		Role:     "customer",
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

func (customerService *CustomerServiceImpl) ViewOne(customerId int, customerName, customerNoHp string) (webcustomer.CustomerResponse, error) {
	customer, err := customerService.CustomerRepository.FindByIdOrNameOrHp(customerId, customerName, customerNoHp)
	if err != nil {
		return webcustomer.CustomerResponse{}, err
	}

	customerResponse := webcustomer.CustomerResponse{
		Id:       customer.Id,
		Name:     customer.Name,
		NoHp:     customer.NoHp,
		Address:  customer.Address,
		Password: customer.Password,
		Role:     customer.Role,
	}

	return customerResponse, nil
}

func (customerService *CustomerServiceImpl) Edit(req webcustomer.CustomerUpdateRequest) (webcustomer.CustomerResponse, error) {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	customer := entity.Customer{
		Id:       req.Id,
		Name:     req.Name,
		NoHp:     req.NoHp,
		Address:  req.Address,
		Password: string(encryptedPassword),
	}

	customerData, err := customerService.CustomerRepository.Update(&customer)
	helper.PanicErr(err)

	customerResponse := webcustomer.CustomerResponse{
		Id:       customerData.Id,
		Name:     customerData.Name,
		NoHp:     customerData.NoHp,
		Address:  customerData.Address,
		Password: customerData.Password,
		Role:     "customer",
		Token:    "NULL",
	}

	return customerResponse, nil
}

func (customerService *CustomerServiceImpl) Unreg(customerId int, customerName, customerNoHp string) (webcustomer.CustomerResponse, error) {
	fmt.Printf("customerId: %v\n", customerId)
	fmt.Scanln()
	customerData, err := customerService.CustomerRepository.FindByIdOrNameOrHp(customerId, customerName, customerNoHp)
	helper.PanicErr(err)

	err = customerService.CustomerRepository.Delete(customerData.Id)
	helper.PanicErr(err)

	customerResponse := webcustomer.CustomerResponse{
		Name:     customerData.Name,
		NoHp:     customerData.NoHp,
		Address:  customerData.Address,
		Password: customerData.Password,
		Role:     customerData.Role,
		Token:    "NULL",
	}

	return customerResponse, nil
}
func (customerService *CustomerServiceImpl) Notification(customerId int) ([]webtransaction.TransactionResponse, error) {
	tc, _ := customerService.CustomerRepository.FindTransactionCustomerById(customerId, 0)

	customerResponse := make([]webtransaction.TransactionResponse, len(tc))
	for i, customer := range tc {
		customerResponse[i] = webtransaction.TransactionResponse{
			CustomerName:   customer.CustomerName,
			MerkName:       customer.MerkName,
			ProductName:    customer.ProductName,
			ProductPrice:   customer.ProductPrice,
			ShippingCost:   customer.ShippingCost,
			Qty:            customer.Qty,
			Tax:            customer.Tax,
			TotalPrice:     int(customer.TotalPrice),
			BuyDate:        customer.BuyDate.Format("2006-01-02"),
			Status:         customer.Status,
			StoreName:      customer.StoreName,
			VirtualAccount: customer.VirtualAccount,
		}
	}

	// tc, _ := customerService.CustomerRepository.FindTransactionCustomerById(customerId, 0)

	// // Map untuk mengumpulkan transaksi berdasarkan virtual account
	// transactionMap := make(map[int][]webtransaction.ProductResponse)
	// totalPrice := 0

	// for _, customer := range tc {
	// 	product := webtransaction.ProductResponse{
	// 		MerkName:     customer.MerkName,
	// 		ProductName:  customer.ProductName,
	// 		ProductPrice: customer.ProductPrice,
	// 		Qty:          customer.Qty,
	// 		Total:        float64(customer.ProductPrice * customer.Qty),
	// 	}

	// 	transactionMap[customer.VirtualAccount] = append(transactionMap[customer.VirtualAccount], product)
	// 	totalPrice += customer.ProductPrice * customer.Qty
	// }

	// // Buat slice untuk menyimpan hasil akhir
	// customerResponse := []webtransaction.TransactionMultiplerResponse{}
	// total_price := float64(totalPrice) + float64(tc[0].ShippingCost) + float64(tc[0].Tax)

	// // Iterasi melalui map dan tambahkan transaksi ke customerResponse
	// for virtualAccount, products := range transactionMap {
	// 	transactionResponse := webtransaction.TransactionMultiplerResponse{
	// 		CustomerName:   tc[0].CustomerName,
	// 		Product:        products,
	// 		ShippingCost:   tc[0].ShippingCost,
	// 		Tax:            tc[0].Tax,
	// 		TotalPrice:     int(total_price),
	// 		BuyDate:        tc[0].BuyDate.Format("2006-01-02"),
	// 		Status:         tc[0].Status,
	// 		StoreName:      tc[0].StoreName,
	// 		VirtualAccount: virtualAccount,
	// 	}

	// 	customerResponse = append(customerResponse, transactionResponse)
	// }

	// // Update total_price pada setiap transaksi di customerResponse
	// for i := range customerResponse {
	// 	customerResponse[i].TotalPrice = int(total_price)
	// }
	return customerResponse, nil
}
