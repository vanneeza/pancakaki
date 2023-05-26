package transactionservice

import (
	webcustomer "pancakaki/internal/domain/web/customer"
	webtransaction "pancakaki/internal/domain/web/transaction"
)

type TransactionService interface {
	MakeOrder(req webtransaction.TransactionOrderCreateRequest) (webtransaction.TransactionResponse, error)
	CustomerPayment(req webtransaction.PaymentCreateRequest) ([]webcustomer.TransactionCustomer, error)
	// MakeDetailOrder(req webtransaction.TransactionOrderCreateRequest) (webtransaction.TransactionResponse, error)
}
