package transactionservice

import (
	"fmt"
	"log"
	"pancakaki/internal/domain/entity"
	webtransaction "pancakaki/internal/domain/web/transaction"
	chartrepository "pancakaki/internal/repository/chart"
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
	ChartRepository       chartrepository.ChartRepository
}

func NewTransactionService(transactionRepository transactionrepository.TransactionRepository,
	productRepository productrepository.ProductRepository, customerRepository customerrepository.CustomerRepository,
	ownerrepository ownerrepository.OwnerRepository, chartRepository chartrepository.ChartRepository) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		ProductRepository:     productRepository,
		CustomerRepository:    customerRepository,
		Ownerrepository:       ownerrepository,
		ChartRepository:       chartRepository,
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
	customer, _ := transactionService.CustomerRepository.FindByIdOrNameOrHp(req.CustomerId, "", "")
	storeName, tax, _ := transactionService.Ownerrepository.GetTaxAndStoreOwner(req.ProductId)
	merkName, _ := transactionService.TransactionRepository.GetMerkNameByProduct(product.Id)

	// TotalPrice : ( QTY * harga_product ) / tax %
	total := product.Price * req.Qty

	taxProduct := float64(req.Qty*product.Price) * (tax / 100)
	totalPrice := total + int(taxProduct) + product.ShippingCost

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

	transactionOrder := []*entity.TransactionOrder{
		{
			Qty:           req.Qty,
			Total:         totalPrice,
			CustomerId:    req.CustomerId,
			ProductId:     req.ProductId,
			DetailOrderId: txDetail.Id,
		},
	}

	transactionData, _ := transactionService.TransactionRepository.CreateOrder(transactionOrder)

	transactionResponse := webtransaction.TransactionResponse{
		CustomerName:   customer.Name,
		MerkName:       merkName,
		ProductName:    product.Name,
		ProductPrice:   product.Price,
		ShippingCost:   product.ShippingCost,
		Qty:            transactionData[0].Qty,
		Tax:            taxProduct,
		TotalPrice:     txDetail.TotalPrice,
		BuyDate:        txDetail.BuyDate.Format("2006-01-02"),
		Status:         "waiting payment",
		StoreName:      storeName,
		VirtualAccount: virtualAccount,
	}
	return transactionResponse, nil
}

func (transactionService *TransactionServiceImpl) CustomerPayment(req webtransaction.PaymentCreateRequest) (webtransaction.CustomerPaymentResponse, error) {
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

	var (
		virtualAccount, qty, productId int
		totalPrice                     float64
		productResponses               []webtransaction.ProductResponse
		txCustomerResponse             entity.TransactionCustomer
	)

	txCustomerData, _ := transactionService.CustomerRepository.FindTransactionCustomerById(0, req.VirtualAccount)

	for _, txCustomer := range txCustomerData {
		productResponse := webtransaction.ProductResponse{
			MerkName:     txCustomer.MerkName,
			ProductName:  txCustomer.ProductName,
			ProductPrice: txCustomer.ProductPrice,
			Qty:          txCustomer.Qty,
		}

		txCustomerResponse = entity.TransactionCustomer{
			CustomerName:   txCustomer.CustomerName,
			ShippingCost:   txCustomer.ShippingCost,
			Tax:            txCustomer.Tax,
			TotalPrice:     txCustomer.TotalPrice,
			Status:         "Paid",
			StoreName:      txCustomer.StoreName,
			BuyDate:        txCustomer.BuyDate,
			VirtualAccount: txCustomer.VirtualAccount,
		}

		productResponses = append(productResponses, productResponse)

		totalPrice = txCustomer.TotalPrice
		virtualAccount = txCustomer.VirtualAccount
		qty = txCustomer.Qty
		productId = txCustomer.ProductId

	}

	if float64(req.Pay) < totalPrice {
		return webtransaction.CustomerPaymentResponse{}, fmt.Errorf("uang kurang")
	}

	if virtualAccount != req.VirtualAccount {
		return webtransaction.CustomerPaymentResponse{}, fmt.Errorf("virtual account salah")
	}

	transactionService.TransactionRepository.CustomerPayment(&pay)
	transactionService.TransactionRepository.UpdatePhotoAndStatus(&txDetail)

	product, _ := transactionService.ProductRepository.FindProductById(productId)

	productStock := product.Stock - qty
	stock := entity.Product{
		Id:    productId,
		Stock: productStock,
	}

	p, _ := transactionService.ProductRepository.UpdateProductStock(&stock)
	fmt.Printf("p: %v\n", p)

	transactionCustomerResponse := webtransaction.CustomerPaymentResponse{
		CustomerName:   txCustomerResponse.CustomerName,
		Product:        productResponses,
		ShippingCost:   txCustomerResponse.ShippingCost,
		Tax:            txCustomerResponse.Tax,
		TotalPrice:     txCustomerResponse.TotalPrice,
		BuyDate:        txCustomerResponse.BuyDate.Format("2006-01-02"),
		Status:         "Paid",
		StoreName:      txCustomerResponse.StoreName,
		VirtualAccount: txCustomerResponse.VirtualAccount,
		Photo:          req.Photo.Filename,
	}

	return transactionCustomerResponse, nil
}

func (transactionService *TransactionServiceImpl) MakeMultipleOrder(req webtransaction.TransactionOrderCreateRequest) (webtransaction.TransactionMultiplerResponse, error) {
	// bth chart repo buat ambil semua datanya

	// Butuh Repository FIND Product By Id
	//---- Buat dapetin harga product dan nama

	// Butuh Repository FIND Customer By Id
	//----- Buat dapetin nama

	// Butuh Repository FIND Owner By Id
	//------ Buat dapetin TAX
	fmt.Printf("customer: %v\n", req.CustomerId)
	fmt.Scanln()
	virtualAccount := helper.GenerateRandomNumber()
	c, _ := transactionService.ChartRepository.FindAll(req.CustomerId)

	var (
		total, tax        float64
		storeName         string
		shippingCost      int
		charts            []entity.Chart
		transactionOrders []*entity.TransactionOrder
		productResponses  []webtransaction.ProductResponse
	)

	for _, chart := range c {
		fmt.Printf("chartTotal: %v\n", chart.Total)
		total = total + chart.Total
		product, _ := transactionService.ProductRepository.FindProductById(chart.ProductId)
		sn, t, _ := transactionService.Ownerrepository.GetTaxAndStoreOwner(chart.ProductId)
		mn, _ := transactionService.TransactionRepository.GetMerkNameByProduct(product.Id)
		charts = append(charts, chart)

		storeName = sn
		tax = t
		shippingCost = product.ShippingCost

		productResponse := webtransaction.ProductResponse{
			MerkName:     mn,
			ProductName:  product.Name,
			ProductPrice: product.Price,
			Qty:          chart.Qty,
		}

		productResponses = append(productResponses, productResponse)
	}

	customer, _ := transactionService.CustomerRepository.FindByIdOrNameOrHp(req.CustomerId, "", "")
	fmt.Printf("total: %v\n", total)
	fmt.Printf("tax: %v\n", tax)
	taxProduct := total * (tax / 100)
	totalPrice := total + taxProduct + float64(shippingCost)

	transactionDetail := entity.TransactionOrderDetail{
		BuyDate:        time.Now().Truncate(24 * time.Hour),
		Status:         "waiting payment",
		TotalPrice:     int(totalPrice),
		Photo:          "None",
		Tax:            taxProduct,
		VirtualAccount: int64(virtualAccount),
	}

	txDetail, _ := transactionService.TransactionRepository.CreateOrderDetail(&transactionDetail)

	for _, c := range charts {
		transactionOrder := entity.TransactionOrder{
			Qty:           c.Qty,
			Total:         int(c.Total),
			CustomerId:    c.CustomerId,
			ProductId:     c.ProductId,
			DetailOrderId: txDetail.Id,
		}

		transactionOrders = append(transactionOrders, &transactionOrder)
	}

	log.Println(transactionOrders, "transactionOrders")
	fmt.Scanln()

	transactionService.TransactionRepository.CreateOrder(transactionOrders)

	transactionResponse := webtransaction.TransactionMultiplerResponse{
		CustomerName:   customer.Name,
		Product:        productResponses,
		ShippingCost:   shippingCost,
		Tax:            taxProduct,
		TotalPrice:     txDetail.TotalPrice,
		BuyDate:        txDetail.BuyDate.Format("2006-01-02"),
		Status:         "waiting payment",
		StoreName:      storeName,
		VirtualAccount: virtualAccount,
	}
	return transactionResponse, nil
}
