package transactionservice

import (
	"fmt"
	"pancakaki/internal/domain/entity"
	webcustomer "pancakaki/internal/domain/web/customer"
	webtransaction "pancakaki/internal/domain/web/transaction"
	customerrepository "pancakaki/internal/repository/customer"
	ownerrepository "pancakaki/internal/repository/owner"
	productrepository "pancakaki/internal/repository/product"
	transactionrepository "pancakaki/internal/repository/transaction"
	"pancakaki/utils/helper"
	"time"
)

type TransactionServiceImpl struct {
	TransactionRepository transactionrepository.TransactionRepository
	ProductRepository     productrepository.ProductRepository
	CustomerRepository    customerrepository.CustomerRepository
	Ownerrepository       ownerrepository.OwnerRepository
}

func NewTransactionService(transactionRepository transactionrepository.TransactionRepository,
	productRepository productrepository.ProductRepository, customerRepository customerrepository.CustomerRepository,
	ownerrepository ownerrepository.OwnerRepository) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		ProductRepository:     productRepository,
		CustomerRepository:    customerRepository,
		Ownerrepository:       ownerrepository,
	}
}

func (transactionService *TransactionServiceImpl) MakeOrder(req webtransaction.TransactionOrderCreateRequest) (webtransaction.TransactionResponse, error) {
	// Butuh Repository FIND Product By Id
	//---- Buat dapetin harga product dan nama

	// Butuh Repository FIND Customer By Id
	//----- Buat dapetin nama

	// Butuh Repository FIND Owner By Id
	//------ Buat dapetin TAX

	virtualAccount := helper.GenerateRandomNumber()

	product, _ := transactionService.ProductRepository.FindProductById(req.ProductId)
	customer, _ := transactionService.CustomerRepository.FindById(req.CustomerId)
	storeName, tax, _ := transactionService.Ownerrepository.GetTaxAndStoreOwner(req.ProductId)
	merkName, _ := transactionService.TransactionRepository.GetMerkNameByProduct(product.Id)

	// TotalPrice : ( QTY * harga_product ) / tax %
	taxProduct := float64(req.Qty*product.Price) * (tax / 100)
	totalPrice := product.Price + int(taxProduct) + product.ShippingCost

	transactionDetail := entity.TransactionOrderDetail{
		BuyDate:        time.Now().Truncate(24 * time.Hour),
		Status:         "waiting payment",
		TotalPrice:     totalPrice,
		Photo:          "None",
		Tax:            taxProduct,
		VirtualAccount: int64(virtualAccount),
	}
	fmt.Printf("transactionDetail: %v\n", transactionDetail)
	fmt.Scanln()

	txDetail, _ := transactionService.TransactionRepository.CreateOrderDetail(&transactionDetail)

	transactionOrder := entity.TransactionOrder{
		Qty:           req.Qty,
		Total:         totalPrice,
		CustomerId:    req.CustomerId,
		ProductId:     req.ProductId,
		DetailOrderId: txDetail.Id,
	}

	transactionData, _ := transactionService.TransactionRepository.CreateOrder(&transactionOrder)

	transactionResponse := webtransaction.TransactionResponse{
		CustomerName:   customer.Name,
		MerkName:       merkName,
		ProductName:    product.Name,
		ProductPrice:   product.Price,
		ShippingCost:   product.ShippingCost,
		Qty:            transactionData.Qty,
		Tax:            taxProduct,
		TotalPrice:     txDetail.TotalPrice,
		BuyDate:        txDetail.BuyDate.Format("2006-01-02"),
		Status:         "waiting payment",
		StoreName:      storeName,
		VirtualAccount: virtualAccount,
	}
	return transactionResponse, nil
}

func (transactionService *TransactionServiceImpl) CustomerPayment(req webtransaction.PaymentCreateRequest) ([]webcustomer.TransactionCustomer, error) {
	// Ambil dulu total price & virtual accountnya
	// Validasi total price, va,  ,wajib photo bukti pembayaran
	// kalo ga sama, return error
	// kalo sama,
	// insert ke payment
	// update status transaction_order_detail + masukin photo

	pay := entity.Payment{
		TransactionDetailOrderId: req.Transaction_detail_order_Id,
		Pay:                      float64(req.Pay),
	}

	txDetail := entity.TransactionOrderDetail{
		Id:     req.Transaction_detail_order_Id,
		Status: "Paid",
		Photo:  req.Photo.Filename,
	}

	var virtualAccount, qty, productId int
	var totalPrice float64

	txCustomerData, _ := transactionService.CustomerRepository.FindTransactionCustomerById(0, req.VirtualAccount)

	transactionCustomerResponse := make([]webcustomer.TransactionCustomer, len(txCustomerData))
	for i, txCustomer := range txCustomerData {
		transactionCustomerResponse[i] = webcustomer.TransactionCustomer{
			CustomerName:   txCustomer.CustomerName,
			MerkName:       txCustomer.MerkName,
			ProductId:      txCustomer.ProductId,
			ProductName:    txCustomer.ProductName,
			ProductPrice:   txCustomer.ProductPrice,
			ShippingCost:   txCustomer.ShippingCost,
			Qty:            txCustomer.Qty,
			Tax:            txCustomer.Tax,
			TotalPrice:     txCustomer.TotalPrice,
			BuyDate:        txCustomer.BuyDate.Format("2006-01-02"),
			Status:         "Paid",
			StoreName:      txCustomer.StoreName,
			VirtualAccount: txCustomer.VirtualAccount,
			Photo:          req.Photo.Filename,
		}
		totalPrice = txCustomer.TotalPrice
		virtualAccount = txCustomer.VirtualAccount
		qty = txCustomer.Qty
		productId = txCustomer.ProductId
	}

	fmt.Printf("txCustomerData: %v\n", txCustomerData)

	fmt.Printf("productId: %v\n", productId)
	fmt.Printf("virtualAccount: %v\n", virtualAccount)
	fmt.Printf("req.VirtualAccount: %v\n", req.VirtualAccount)
	fmt.Scanln()
	if float64(req.Pay) < totalPrice {
		return []webcustomer.TransactionCustomer{}, fmt.Errorf("uang kurang")
	}

	if virtualAccount != req.VirtualAccount {
		return []webcustomer.TransactionCustomer{}, fmt.Errorf("virtual account salah")
	}

	transactionService.TransactionRepository.CustomerPayment(&pay)
	transactionService.TransactionRepository.UpdatePhotoAndStatus(&txDetail)

	product, _ := transactionService.ProductRepository.FindProductById(productId)
	// Update Stock ( - qty )

	productStock := product.Stock - int16(qty)
	stock := entity.Product{
		Id:    productId,
		Stock: productStock,
	}
	fmt.Printf("stock: %v\n", stock)

	p, _ := transactionService.ProductRepository.UpdateProductStock(&stock)
	fmt.Printf("p: %v\n", p)

	return transactionCustomerResponse, nil
}
