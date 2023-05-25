package paymentrepository

import (
	"database/sql"
	"fmt"
	"pancakaki/internal/domain/entity"
)

type PaymentRepository interface {
	InsertPayment(newPayment *entity.Payment) (*entity.Payment, error)
	// UpdatePayment(updatePayment *entity.Payment) (*entity.Payment, error)
	// DeletePayment(deletePayment *entity.Payment) error
	// FindPaymentById(id int) (*entity.Payment, error)
	// FindPaymentByName(name string) (*entity.Payment, error)
	// FindAllPayment() ([]entity.Payment, error)
}

type paymentRepository struct {
	db *sql.DB
	// storeRepo
}

// InsertPayment implements PaymentRepository
func (repo *paymentRepository) InsertPayment(newPayment *entity.Payment) (*entity.Payment, error) {

	stmt, err := repo.db.Prepare("INSERT INTO tbl_payment (name,qty,price) VALUES ($1,$2,$3) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to insert payment : %w", err)
	}
	defer stmt.Close()

	// createdAt := time.Now()
	err = stmt.QueryRow(newPayment.Name, newPayment.Qty, newPayment.Price).Scan(&newPayment.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert payment : %w", err)
	}

	return newPayment, nil
}

// UpdatePayment implements PaymentRepository
// func (repo *paymentRepository) UpdatePayment(updatePayment *entity.Payment) (*entity.Payment, error) {
// 	stmt, err := repo.db.Prepare("UPDATE tbl_payment SET name=$1, price=$2, stock=$3, description=$4, shipping_cost=$5,merk_id=$6,store_id=$7 WHERE id = $8")
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to update payment : %w", err)
// 	}
// 	defer stmt.Close()

// 	// updateAt := time.Now()
// 	_, err = stmt.Exec(updatePayment.Id, updatePayment.Name, updatePayment.Price, updatePayment.Stock, updatePayment.Description, updatePayment.ShippingCost, updatePayment.MerkId, updatePayment.StoreId)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to update payment : %w", err)
// 	}

// 	return updatePayment, nil
// }

// func (repo *paymentRepository) DeletePayment(deletePayment *entity.Payment) error {
// 	stmt, err := repo.db.Prepare("UPDATE tbl_payment SET is_delete = true WHERE id = $1")
// 	if err != nil {
// 		return fmt.Errorf("failed to delete payment : %w", err)
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(deletePayment.Id)
// 	if err != nil {
// 		return fmt.Errorf("failed to delete payment : %w", err)
// 	}

// 	return nil
// }

// // FindAllPayment implements PaymentRepository
// func (repo *paymentRepository) FindAllPayment() ([]entity.Payment, error) {
// 	var payments []entity.Payment
// 	rows, err := repo.db.Query("SELECT id,name,price,stock,description,shipping_cost,merk_id,store_id FROM tbl_payment where store_id = $1")
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get payment : %w", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var payment entity.Payment
// 		err := rows.Scan(&payment.Id, &payment.Name, &payment.Price, &payment.Stock, &payment.Description, &payment.ShippingCost, &payment.MerkId, &payment.StoreId)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to get payment : %w", err)
// 		}
// 		payments = append(payments, payment)
// 	}

// 	return payments, nil
// }

// // FindPaymentById implements PaymentRepository
// func (repo *paymentRepository) FindPaymentById(id int) (*entity.Payment, error) {
// 	var payment entity.Payment
// 	stmt, err := repo.db.Prepare("SELECT id,name,price,stock,description,shipping_cost,merk_id,store_id FROM tbl_payment WHERE id = $1")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	err = stmt.QueryRow(id).Scan(&payment.Id, &payment.Name, &payment.Price, &payment.Stock, &payment.Description, &payment.ShippingCost, &payment.MerkId, &payment.StoreId)
// 	if err == sql.ErrNoRows {
// 		return nil, fmt.Errorf("payment with id %d not found", id)
// 	} else if err != nil {
// 		return nil, err
// 	}

// 	return &payment, nil
// }

// // FindPaymentByName implements PaymentRepository
// func (repo *paymentRepository) FindPaymentByName(name string) (*entity.Payment, error) {
// 	var payment entity.Payment
// 	stmt, err := repo.db.Prepare("SELECT id,name,price,stock,description,shipping_cost,merk_id,store_id FROM tbl_payment WHERE name = $1")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	err = stmt.QueryRow(name).Scan(&payment.Id, &payment.Name, &payment.Price, &payment.Stock, &payment.Description, &payment.ShippingCost, &payment.MerkId, &payment.StoreId)
// 	if err == sql.ErrNoRows {
// 		return nil, fmt.Errorf("payment with name %s not found", name)
// 	} else if err != nil {
// 		return nil, err
// 	}

// 	return &payment, nil
// }

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{db: db}
}
