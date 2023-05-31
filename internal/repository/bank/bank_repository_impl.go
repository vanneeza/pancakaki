package bankrepository

import (
	"database/sql"
	"pancakaki/internal/domain/entity"
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

	stmt, err := r.Db.Prepare("INSERT INTO tbl_bank (name, bank_account, account_name) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(bank.Name, bank.BankAccount, bank.AccountName).Scan(&bank.Id)
	if err != nil {
		return nil, err
	}

	return bank, nil
}

func (r *BankRepositoryImpl) CreateBankAdmin(bankAdmin *entity.BankAdmin) (*entity.BankAdmin, error) {
	stmt, err := r.Db.Prepare("INSERT INTO tbl_bank_admin (admin_id, bank_id) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(bankAdmin.AdminId, bankAdmin.BankId).Scan(&bankAdmin.Id)
	if err != nil {
		return nil, err
	}

	return bankAdmin, nil
}

func (r *BankRepositoryImpl) FindAll() ([]entity.Bank, error) {
	var tbl_bank []entity.Bank
	rows, err := r.Db.Query(`SELECT tbl_bank.id, tbl_bank.name, tbl_bank.bank_account, tbl_bank.account_name
	FROM tbl_bank INNER JOIN tbl_bank_admin ON tbl_bank.id = tbl_bank_admin.bank_id where tbl_bank.is_deleted = false`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bank entity.Bank
		err := rows.Scan(&bank.Id, &bank.Name, &bank.BankAccount, &bank.AccountName)
		if err != nil {
			return nil, err
		}
		tbl_bank = append(tbl_bank, bank)
	}

	return tbl_bank, nil
}

func (r *BankRepositoryImpl) FindById(bankId int) ([]entity.Bank, error) {

	var tbl_bank []entity.Bank
	rows, err := r.Db.Query(`SELECT tbl_bank.id, tbl_bank.name, tbl_bank.bank_account, tbl_bank.account_name
	FROM tbl_bank INNER JOIN tbl_bank_admin ON tbl_bank.id = tbl_bank_admin.bank_id where tbl_bank.is_deleted = false AND tbl_bank.id = $1`, bankId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bank entity.Bank
		err := rows.Scan(&bank.Id, &bank.Name, &bank.BankAccount, &bank.AccountName)
		if err != nil {
			return nil, err
		}
		tbl_bank = append(tbl_bank, bank)
	}

	return tbl_bank, nil
}

func (r *BankRepositoryImpl) Update(bank *entity.Bank) (*entity.Bank, error) {
	stmt, err := r.Db.Prepare(`UPDATE tbl_bank SET name = $1, bank_account = $2, account_name = $3	WHERE id = $4`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(bank.Name, bank.BankAccount, bank.AccountName, bank.Id)
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
