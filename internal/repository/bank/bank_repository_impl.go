package bankrepository

import (
	"database/sql"
	entity "pancakaki/internal/domain/entity/bank"
)

type BankRepositoryImpl struct {
	Db *sql.DB
}

func NewBankRepository(Db *sql.DB) BankRepository {
	return &BankRepositoryImpl{
		Db: Db,
	}
}

func (r *BankRepositoryImpl) Create(bank *entity.Bank) (*entity.Bank, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_bank (name, bank_account) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(bank.Name, bank.BankAccount).Scan(&bank.Id)
	if err != nil {
		return nil, err
	}

	return bank, nil
}

func (r *BankRepositoryImpl) FindAll() ([]entity.Bank, error) {
	var tbl_bank []entity.Bank
	rows, err := r.Db.Query("SELECT id, name, bank_account FROM tbl_bank WHERE is_deleted FALSE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bank entity.Bank
		err := rows.Scan(&bank.Id, &bank.Name, &bank.BankAccount)
		if err != nil {
			return nil, err
		}
		tbl_bank = append(tbl_bank, bank)
	}

	return tbl_bank, nil
}

func (r *BankRepositoryImpl) FindById(id int) (*entity.Bank, error) {
	var bank entity.Bank
	stmt, err := r.Db.Prepare("SELECT id, name, bank_account FROM tbl_bank WHERE id = $1 AND is_deleted = FALSE")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&bank.Id, &bank.Name, &bank.BankAccount)
	if err != nil {
		return nil, err
	}

	return &bank, nil
}

func (r *BankRepositoryImpl) Update(bank *entity.Bank) (*entity.Bank, error) {
	stmt, err := r.Db.Prepare("UPDATE tbl_bank SET name = $1, bank_account = $2 WHERE id = $3 AND is_deleted = FALSE")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(bank.Name, bank.BankAccount, bank.Id)
	if err != nil {
		return nil, err
	}

	return bank, nil
}

func (r *BankRepositoryImpl) Delete(bankId int) error {
	stmt, err := r.Db.Prepare("UPDATE tbl_bank SET is_deleted = TRUE WHERE id= $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(bankId)
	if err != nil {
		return err
	}

	return nil
}
