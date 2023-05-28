package bankstorerepository

import (
	"database/sql"
	"fmt"
	"pancakaki/internal/domain/entity"
)

type BankStoreRepository interface {
	CreateBankStore(newBank *entity.BankStore, tx *sql.Tx) (*entity.BankStore, error)
	CreateBank(newBank *entity.Bank, tx *sql.Tx) (*entity.Bank, error)
	UpdateBankStore(updateBank *entity.Bank, tx *sql.Tx) (*entity.Bank, error)
	DeleteBank(id int, tx *sql.Tx) error
	DeleteBankStore(storeId int, tx *sql.Tx) error
	GetBankStoreByStoreId(id int) ([]entity.Bank, error)
	GetBankAdminById(id int) ([]entity.Bank, error)
}

type bankStoreRepository struct {
	db *sql.DB
}

func (repo *bankStoreRepository) GetBankAdminById(id int) ([]entity.Bank, error) {
	var banks []entity.Bank
	rows, err := repo.db.Query(`
		SELECT tbl_bank.id, tbl_bank.name, tbl_bank.bank_account, tbl_bank.account_name FROM tbl_bank
		INNER JOIN tbl_bank_admin ON tbl_bank.id = tbl_bank_admin.bank_id WHERE tbl_bank_admin.admin_id = $1`, id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("bank with admin id %d not found", id)
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var bank entity.Bank
		err := rows.Scan(&bank.Id, &bank.Name, &bank.BankAccount, &bank.AccountName)
		if err != nil {
			return nil, fmt.Errorf("failed to get bank admin : %w", err)
		}
		banks = append(banks, bank)
	}

	return banks, nil
}

func (repo *bankStoreRepository) GetBankStoreByStoreId(id int) ([]entity.Bank, error) {
	var banks []entity.Bank
	rows, err := repo.db.Query(`
		SELECT tbl_bank.id, tbl_bank.name, tbl_bank.bank_account, tbl_bank.account_name FROM tbl_bank
		INNER JOIN tbl_bank_store ON tbl_bank.id = tbl_bank_store.bank_id WHERE tbl_bank_store.store_id = $1`, id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("bank with store id %d not found", id)
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var bank entity.Bank
		err := rows.Scan(&bank.Id, &bank.Name, &bank.BankAccount, &bank.AccountName)
		if err != nil {
			return nil, fmt.Errorf("failed to get bank store: %w", err)
		}
		banks = append(banks, bank)
	}

	return banks, nil
}

func (repo *bankStoreRepository) CreateBank(newBank *entity.Bank, tx *sql.Tx) (*entity.Bank, error) {
	stmt, err := repo.db.Prepare("INSERT INTO tbl_bank (name,bank_account,account_name) VALUES ($1,$2,$3) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to create bank store: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newBank.Name, newBank.BankAccount, newBank.AccountName).Scan(&newBank.Id)
	bankValidate(err, "create bank store", tx)
	return newBank, nil
}

func (repo *bankStoreRepository) CreateBankStore(newBankStore *entity.BankStore, tx *sql.Tx) (*entity.BankStore, error) {
	stmt, err := repo.db.Prepare("INSERT INTO tbl_bank_store (store_id,bank_id) VALUES ($1,$2) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to create bank store: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newBankStore.StoreId, newBankStore.BankId).Scan(&newBankStore.Id)
	bankValidate(err, "create bank store", tx)
	return newBankStore, nil
}

func (repo *bankStoreRepository) UpdateBankStore(updateBank *entity.Bank, tx *sql.Tx) (*entity.Bank, error) {
	stmt, err := repo.db.Prepare("UPDATE tbl_bank SET name = $1, bank_account=$2, account_name=$3 WHERE id = $4")
	if err != nil {
		return nil, fmt.Errorf("failed to update bank store: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateBank.Name, updateBank.BankAccount, updateBank.AccountName, updateBank.Id)
	bankValidate(err, "update bank store", tx)

	return updateBank, nil
}

func (repo *bankStoreRepository) DeleteBank(id int, tx *sql.Tx) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_bank SET is_deleted = true WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to delete bank store: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	bankValidate(err, "delete bank store", tx)

	return nil
}

func (repo *bankStoreRepository) DeleteBankStore(storeId int, tx *sql.Tx) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_bank_store SET is_deleted = true WHERE store_id = $1")
	if err != nil {
		return fmt.Errorf("failed to delete bank store: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(storeId)
	bankValidate(err, "delete bank store", tx)

	return nil
}

func bankValidate(err error, message string, tx *sql.Tx) {
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
