package storerepository

import (
	"database/sql"
	"fmt"
	entitybank "pancakaki/internal/domain/entity"
	entitybankstore "pancakaki/internal/domain/entity/bank_store"
	entitystore "pancakaki/internal/domain/entity/store"
	webstore "pancakaki/internal/domain/web/store"

	// bankrepository "pancakaki/internal/repository/bank"
	bankstorerepository "pancakaki/internal/repository/bank"
)

type StoreRepository interface {
	GetStoreByOwnerId(id int) (*entitystore.Store, error)
	GetStoreByName(name string) (*entitystore.Store, error)
	CreateStore(newStore *entitystore.Store, tx *sql.Tx) (*entitystore.Store, error)
	CreateMainStore(newTransactionStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error)
	UpdateStore(updateStore *entitystore.Store, tx *sql.Tx) (*entitystore.Store, error)
	UpdateMainStore(newUpdateStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error)
}

type storeRepository struct {
	db *sql.DB
	// repoBank      bankstorerepository.BankStoreRepository
	repoBankStore bankstorerepository.BankStoreRepository
}

func (repo *storeRepository) CreateMainStore(newTransactionStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		panic(err)
	}
	newStore := entitystore.Store{
		Name:    newTransactionStore.Name,
		NoHp:    newTransactionStore.NoHp,
		Email:   newTransactionStore.Email,
		Address: newTransactionStore.Address,
		OwnerId: newTransactionStore.OwnerId,
	}
	newBank := entitybank.Bank{
		Name:        newTransactionStore.BankName,
		BankAccount: newTransactionStore.BankAccount,
		AccountName: newTransactionStore.AccountName,
	}

	store, err := repo.CreateStore(&newStore, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create store : %w", err)
	}

	bank, err := repo.repoBankStore.CreateBank(&newBank, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create bank : %w", err)
	}

	newBankStore := entitybankstore.BankStore{
		StoreId: store.Id,
		BankId:  bank.Id,
	}

	_, err = repo.repoBankStore.CreateBankStore(&newBankStore, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create bank store : %w", err)
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		return nil, fmt.Errorf("failed to create store : %w", errCommit)
	}

	storeRespose := webstore.StoreCreateResponse{
		Name:        store.Name,
		NoHp:        store.NoHp,
		Email:       store.Email,
		Address:     store.Address,
		BankName:    bank.Name,
		BankAccount: bank.BankAccount,
	}

	return &storeRespose, nil
}

func (repo *storeRepository) CreateStore(newStore *entitystore.Store, tx *sql.Tx) (*entitystore.Store, error) {
	stmt, err := repo.db.Prepare("INSERT INTO tbl_store (name, no_hp, email, address, owner_id) VALUES ($1,$2,$3,$4,$5) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to create store : %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newStore.Name, newStore.NoHp, newStore.Email, newStore.Address, newStore.OwnerId).Scan(&newStore.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create store : %w", err)
	// }
	validate(err, "create store", tx)

	return newStore, nil
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "transaction rollback")
	} else {
		fmt.Println("success")
	}
}

func (repo *storeRepository) GetStoreByOwnerId(id int) (*entitystore.Store, error) {
	var store entitystore.Store
	stmt, err := repo.db.Prepare("SELECT id, name,no_hp,email,address,owner_id FROM tbl_store WHERE owner_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&store.Id, &store.Name, &store.NoHp, &store.Email, &store.Address, &store.OwnerId)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("store with owner_id %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &store, nil
}

func (repo *storeRepository) GetStoreByName(name string) (*entitystore.Store, error) {
	var store entitystore.Store
	stmt, err := repo.db.Prepare("SELECT id, name,no_hp,email,address,owner_id FROM tbl_store WHERE name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&store.Id, &store.Name, &store.NoHp, &store.Email, &store.Address, &store.OwnerId)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("store with name %s not found", name)
	} else if err != nil {
		return nil, err
	}

	return &store, nil
}

func (repo *storeRepository) UpdateStore(updateStore *entitystore.Store, tx *sql.Tx) (*entitystore.Store, error) {
	stmt, err := repo.db.Prepare("UPDATE tbl_store SET name = $1, no_hp=$2,email=$3,address=$4 WHERE id = $5")
	if err != nil {
		return nil, fmt.Errorf("failed to update store : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateStore.Name, updateStore.NoHp, updateStore.Email, updateStore.Address, updateStore.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to update store : %w", err)
	// }
	validate(err, "update store", tx)

	return updateStore, nil
}

func (repo *storeRepository) UpdateMainStore(newUpdateStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		panic(err)
	}
	updateStore := entitystore.Store{
		Name:    newUpdateStore.Name,
		NoHp:    newUpdateStore.NoHp,
		Email:   newUpdateStore.Email,
		Address: newUpdateStore.Address,
		OwnerId: newUpdateStore.OwnerId,
	}
	updateBank := entitybank.Bank{
		Name:        newUpdateStore.BankName,
		BankAccount: newUpdateStore.BankAccount,
		AccountName: newUpdateStore.AccountName,
	}

	store, err := repo.UpdateStore(&updateStore, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to update store : %w", err)
	}

	bank, err := repo.repoBankStore.UpdateBank(&updateBank, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to update bank : %w", err)
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		return nil, fmt.Errorf("failed to update store : %w", errCommit)
	}

	storeRespose := webstore.StoreCreateResponse{
		Name:        store.Name,
		NoHp:        store.NoHp,
		Email:       store.Email,
		Address:     store.Address,
		BankName:    bank.Name,
		BankAccount: bank.BankAccount,
	}

	return &storeRespose, nil
}

func NewStoreRepository(
	db *sql.DB,
	// repoBank bankrepository.BankRepository,
	repoBankStore bankstorerepository.BankStoreRepository,
) StoreRepository {
	return &storeRepository{
		db: db,
		// repoBank:      repoBank,
		repoBankStore: repoBankStore}
}
