package transactionservice

import webtransaction "pancakaki/internal/domain/web/transaction"

type TransactionService interface {
	MakeOrder(req webtransaction.TransactionOrderCreateRequest) (webtransaction.TransactionResponse, error)
	// MakeDetailOrder(req webtransaction.TransactionOrderCreateRequest) (webtransaction.TransactionResponse, error)
}
