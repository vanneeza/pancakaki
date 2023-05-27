package transactionrepository

import "pancakaki/internal/domain/entity"

type TransactionRepository interface {
	CreateOrder(transactionOrder []*entity.TransactionOrder) ([]*entity.TransactionOrder, error)
	CreateOrderDetail(transactionOrder *entity.TransactionOrderDetail) (*entity.TransactionOrderDetail, error)
	GetMerkNameByProduct(productId int) (string, error)
	CustomerPayment(payment *entity.Payment) (*entity.Payment, error)
	UpdatePhotoAndStatus(TransactionOrderDetail *entity.TransactionOrderDetail) (*entity.TransactionOrderDetail, error)
	// FindAll() ([]entity.TransactionOrder, error)
	// FindById(id int) (*entity.TransactionOrder, error)
	// Delete(transactionOrderId int) error
}
