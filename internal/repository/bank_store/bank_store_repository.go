package bankstorerepository

import (
	"database/sql"
	"fmt"
	entity "pancakaki/internal/domain/entity/bank_store"
)

type BankStoreRepository interface {
	CreateBankStore(newBank *entity.BankStore, tx *sql.Tx) (*entity.BankStore, error)
}

type bankStoreRepository struct {
	db *sql.DB
}

func (repo *bankStoreRepository) CreateBankStore(newBankStore *entity.BankStore, tx *sql.Tx) (*entity.BankStore, error) {
	stmt, err := repo.db.Prepare("INSERT INTO tbl_bank_store (store_id,bank_id) VALUES ($1,$2) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to create bank store: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newBankStore.StoreId, newBankStore.BankId).Scan(&newBankStore.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create bank store: %w", err)
	// }
	validate(err, "create bank store", tx)
	return newBankStore, nil
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "transaction rollback")
	} else {
		fmt.Println("success")
	}
}

func NewBankStoreRepository(db *sql.DB) BankStoreRepository {
	return &bankStoreRepository{db: db}
}
