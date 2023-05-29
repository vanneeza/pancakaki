package paymentrepository

import (
	"database/sql"
	"fmt"
	"pancakaki/internal/domain/entity"
)

type PaymentRepository interface {
	InsertPayment(newPayment *entity.Payment) (*entity.Payment, error)
}

type paymentRepository struct {
	db *sql.DB
}

func (repo *paymentRepository) InsertPayment(newPayment *entity.Payment) (*entity.Payment, error) {

	stmt, err := repo.db.Prepare("INSERT INTO tbl_payment (name,qty,price) VALUES ($1,$2,$3) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to insert payment : %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newPayment.Pay, newPayment.TransactionDetailOrderId, newPayment.Pay).Scan(&newPayment.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert payment : %w", err)
	}

	return newPayment, nil
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{db: db}
}
