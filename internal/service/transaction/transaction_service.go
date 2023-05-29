package transactionservice

import (
	webtransaction "pancakaki/internal/domain/web/transaction"
)

type TransactionService interface {
	MakeOrder(req webtransaction.TransactionOrderCreateRequest) (webtransaction.TransactionResponse, error)
	MakeMultipleOrder(webtransaction.TransactionOrderCreateRequest) (webtransaction.TransactionMultiplerResponse, error)
	CustomerPayment(req webtransaction.PaymentCreateRequest) (webtransaction.CustomerPaymentResponse, error)
	// MakeDetailOrder(req webtransaction.TransactionOrderCreateRequest) (webtransaction.TransactionResponse, error)
}
