package transactionrepository

import (
	"database/sql"
	"fmt"
	"log"
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

func (r *TransactionRepositoryImpl) CreateOrder(transactionOrder *entity.TransactionOrder) (*entity.TransactionOrder, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_transaction_order (quantity, total, customer_id, product_id, detail_order_id) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	stmtStr := fmt.Sprintf("INSERT INTO tbl_transaction_order (quantity, total, customer_id, product_id, detail_order_id) VALUES (%d, %d, %d, %d, %d) RETURNING id",
		transactionOrder.Qty, transactionOrder.Total, transactionOrder.CustomerId, transactionOrder.ProductId, transactionOrder.DetailOrderId)
	log.Println("SQL statement:", stmtStr)
	fmt.Scanln()
	err = stmt.QueryRow(transactionOrder.Qty, transactionOrder.Total, transactionOrder.CustomerId, transactionOrder.ProductId, transactionOrder.DetailOrderId).Scan(&transactionOrder.Id)
	if err != nil {
		return nil, err
	}

	return transactionOrder, nil
}

func (r *TransactionRepositoryImpl) CreateOrderDetail(TransactionOrderDetail *entity.TransactionOrderDetail) (*entity.TransactionOrderDetail, error) {
	log.Println(TransactionOrderDetail, "direpo")
	fmt.Scanln()
	stmt, err := r.Db.Prepare(`INSERT INTO tbl_transaction_detail_order (buy_date, status, total_price, photo, tax) VALUES ($1,$2,$3,$4, $5) returning id`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	stmtStr := fmt.Sprintf("INSERT INTO tbl_transaction_detail_order (buy_date, status, total_price, photo, tax) VALUES ('%v', '%s', %d, '%s', %f) returning id",
		TransactionOrderDetail.BuyDate, TransactionOrderDetail.Status, TransactionOrderDetail.TotalPrice, TransactionOrderDetail.Photo, TransactionOrderDetail.Tax)
	log.Println("SQL statement:", stmtStr)
	fmt.Scanln()

	err = stmt.QueryRow(
		TransactionOrderDetail.BuyDate,
		TransactionOrderDetail.Status,
		TransactionOrderDetail.TotalPrice,
		TransactionOrderDetail.Photo,
		TransactionOrderDetail.Tax).Scan(&TransactionOrderDetail.Id)

	if err != nil {
		return nil, err
	}

	return TransactionOrderDetail, nil
}
