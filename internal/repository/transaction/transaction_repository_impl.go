package transactionrepository

import (
	"database/sql"
	"fmt"
	"pancakaki/internal/domain/entity"
)

type TransactionRepositoryImpl struct {
	Db *sql.DB
}

func NewTransactionRepository(Db *sql.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		Db: Db,
	}
}

func (r *TransactionRepositoryImpl) CreateOrder(transactionOrders []*entity.TransactionOrder) ([]*entity.TransactionOrder, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_transaction_order (quantity, total, customer_id, product_id, detail_order_id) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, transactionOrder := range transactionOrders {
		err = stmt.QueryRow(transactionOrder.Qty, transactionOrder.Total, transactionOrder.CustomerId, transactionOrder.ProductId, transactionOrder.DetailOrderId).Scan(&transactionOrder.Id)
		if err != nil {
			return nil, err
		}
	}

	return transactionOrders, nil
}

func (r *TransactionRepositoryImpl) CreateOrderDetail(TransactionOrderDetail *entity.TransactionOrderDetail) (*entity.TransactionOrderDetail, error) {
	stmt, err := r.Db.Prepare(`INSERT INTO tbl_transaction_detail_order (buy_date, status, total_price, photo, tax, virtual_account) VALUES ($1,$2,$3,$4, $5, $6) returning id`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		TransactionOrderDetail.BuyDate,
		TransactionOrderDetail.Status,
		TransactionOrderDetail.TotalPrice,
		TransactionOrderDetail.Photo,
		TransactionOrderDetail.Tax,
		TransactionOrderDetail.VirtualAccount).Scan(&TransactionOrderDetail.Id)

	if err != nil {
		return nil, err
	}

	return TransactionOrderDetail, nil
}

func (r *TransactionRepositoryImpl) GetMerkNameByProduct(productId int) (string, error) {

	var merkName string

	stmt, err := r.Db.Prepare(`SELECT tbl_merk.name FROM tbl_merk
	iNNER JOIN tbl_product ON tbl_merk.id = tbl_product.merk_id WHERE tbl_product.id = $1`)
	if err != nil {
		return merkName, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(productId).Scan(&merkName)
	if err == sql.ErrNoRows {
		return merkName, fmt.Errorf("merk with product id %d not found", productId)
	} else if err != nil {
		return merkName, err
	}

	return merkName, nil

}

func (r *TransactionRepositoryImpl) CustomerPayment(payment *entity.Payment) (*entity.Payment, error) {

	stmt, err := r.Db.Prepare(`INSERT INTO tbl_payment (transaction_detail_order_id, pay) VALUES ($1,$2) returning id`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(payment.TransactionDetailOrderId, payment.Pay).Scan(&payment.Id)

	if err != nil {
		return nil, err
	}

	return payment, nil

}

func (r *TransactionRepositoryImpl) UpdatePhotoAndStatus(TransactionOrderDetail *entity.TransactionOrderDetail) (*entity.TransactionOrderDetail, error) {

	stmt, err := r.Db.Prepare(`UPDATE tbl_transaction_detail_order SET status = $1, photo = $2 WHERE id = $3`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(TransactionOrderDetail.Status, TransactionOrderDetail.Photo, TransactionOrderDetail.Id)
	if err != nil {
		return nil, err
	}

	return TransactionOrderDetail, nil

}

func (r *TransactionRepositoryImpl) FindById(id int) (*entity.TransactionOrderDetail, error) {
	var transactionOrderDetail entity.TransactionOrderDetail
	stmt, err := r.Db.Prepare("SELECT id, buy_date,status,total_price,photo,tax, virtual_account FROM tbl_transaction_detail_order WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&transactionOrderDetail.Id,
		&transactionOrderDetail.BuyDate,
		&transactionOrderDetail.Status,
		&transactionOrderDetail.TotalPrice,
		&transactionOrderDetail.Photo,
		&transactionOrderDetail.Tax,
		&transactionOrderDetail.VirtualAccount)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("transaction order with id %d not found", transactionOrderDetail.Id)
	} else if err != nil {
		return nil, err
	}

	return &transactionOrderDetail, nil
}
