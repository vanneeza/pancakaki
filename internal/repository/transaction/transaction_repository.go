package transactionrepository

import "pancakaki/internal/domain/entity"

type TransactionRepository interface {
	CreateOrder(transactionOrder *entity.TransactionOrder) (*entity.TransactionOrder, error)
	CreateOrderDetail(transactionOrder *entity.TransactionOrderDetail) (*entity.TransactionOrderDetail, error)

	// FindAll() ([]entity.TransactionOrder, error)
	// FindById(id int) (*entity.TransactionOrder, error)
	// Update(transactionOrder *entity.TransactionOrder) (*entity.TransactionOrder, error)
	// Delete(transactionOrderId int) error
}
