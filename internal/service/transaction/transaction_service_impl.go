package transactionservice

import (
	"fmt"
	"log"
	"pancakaki/internal/domain/entity"
	webtransaction "pancakaki/internal/domain/web/transaction"
	transactionrepository "pancakaki/internal/repository/transaction"
	"time"
)

type TransactionServiceImpl struct {
	TransactionRepository transactionrepository.TransactionRepository
}

func NewTransactionService(transactionRepository transactionrepository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
	}
}

func (transactionService *TransactionServiceImpl) MakeOrder(req webtransaction.TransactionOrderCreateRequest) (webtransaction.TransactionResponse, error) {

	// Butuh Repository FIND Product By Id
	//---- Buat dapetin harga product dan nama
	// Butuh Repository FIND Customer By Id
	//----- Buat dapetin nama
	// Butuh Repository FIND Owner By Id
	//------ Buat dapetin TAX

	// TotalPrice : QTY * harga_product
	totalPrice := req.Qty * 2000000

	transactionDetail := entity.TransactionOrderDetail{
		BuyDate:    time.Now().Truncate(24 * time.Hour),
		Status:     "prepared order",
		TotalPrice: int64(totalPrice),
		Photo:      "document/upload/customer/customerName-payment.jpeg",
		Tax:        4,
	}

	log.Println(transactionDetail, "transactionDetail")
	fmt.Scanln()

	txDetail, _ := transactionService.TransactionRepository.CreateOrderDetail(&transactionDetail)

	log.Println(txDetail, "txDetail")
	fmt.Scanln()
	transactionOrder := entity.TransactionOrder{
		Qty:           req.Qty,
		Total:         totalPrice,
		CustomerId:    req.CustomerId,
		ProductId:     req.ProductId,
		DetailOrderId: txDetail.Id,
	}
	log.Println(transactionOrder, "transactionOrder")
	log.Println(transactionOrder.Qty, "transactionOrder")
	fmt.Scanln()
	transactionData, _ := transactionService.TransactionRepository.CreateOrder(&transactionOrder)

	transactionResponse := webtransaction.TransactionResponse{
		ProductName: "S7",
		Qty:         transactionData.Qty,
		Tax:         txDetail.Tax,
		TotalPrice:  txDetail.TotalPrice,
		BuyDate:     txDetail.BuyDate,
	}
	return transactionResponse, nil
}
