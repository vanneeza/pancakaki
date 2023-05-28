package storerepository

import (
	"database/sql"
	"fmt"
	"pancakaki/internal/domain/entity"
	webstore "pancakaki/internal/domain/web/store"
	bankstorerepository "pancakaki/internal/repository/bank_store"
	productrepository "pancakaki/internal/repository/product"
	"strconv"
)

type StoreRepository interface {
	GetStoreByOwnerId(id int) ([]entity.Store, error)
	GetStoreByName(name string) (*entity.Store, error)
	CreateStore(newStore *entity.Store, tx *sql.Tx) (*entity.Store, error)
	CreateMainStore(newTransactionStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error)
	UpdateStore(updateStore *entity.Store, tx *sql.Tx) (*entity.Store, error)
	UpdateMainStore(newUpdateStore *webstore.StoreUpdateRequest) (*webstore.StoreCreateResponse, error)
	DeleteStore(id int, tx *sql.Tx) error
	DeleteMainStore(storeid int, ownerId int) error
}

type storeRepository struct {
	db                  *sql.DB
	bankStoreRepository bankstorerepository.BankStoreRepository
	productRepository   productrepository.ProductRepository
}

func NewStoreRepository(
	db *sql.DB,
	bankStoreRepository bankstorerepository.BankStoreRepository,
	productRepository productrepository.ProductRepository) StoreRepository {
	return &storeRepository{
		db:                  db,
		bankStoreRepository: bankStoreRepository,
		productRepository:   productRepository,
	}
}

func (repo *storeRepository) CreateMainStore(newTransactionStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		panic(err)
	}
	newStoreNoHp, err := strconv.Atoi(newTransactionStore.NoHp)
	if err != nil {
		return nil, err
	}
	newStore := entity.Store{
		Name:    newTransactionStore.Name,
		NoHp:    newStoreNoHp,
		Email:   newTransactionStore.Email,
		Address: newTransactionStore.Address,
		OwnerId: newTransactionStore.OwnerId,
	}
	newBank := entity.Bank{
		Name:        newTransactionStore.BankName,
		BankAccount: newTransactionStore.BankAccount,
		AccountName: newTransactionStore.AccountName,
	}

	store, err := repo.CreateStore(&newStore, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create store : %w", err)
	}

	bank, err := repo.bankStoreRepository.CreateBank(&newBank, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create bank : %w", err)
	}

	newBankStore := entity.BankStore{
		StoreId: store.Id,
		BankId:  bank.Id,
	}

	_, err = repo.bankStoreRepository.CreateBankStore(&newBankStore, tx)
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

func (repo *storeRepository) CreateStore(newStore *entity.Store, tx *sql.Tx) (*entity.Store, error) {
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

func (repo *storeRepository) GetStoreByOwnerId(id int) ([]entity.Store, error) {
	var stores []entity.Store
	rows, err := repo.db.Query("SELECT id, name,no_hp,email,address,owner_id FROM tbl_store WHERE owner_id = $1", id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("store with owner_id %d not found", id)
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var store entity.Store
		err := rows.Scan(&store.Id, &store.Name, &store.NoHp, &store.Email, &store.Address, &store.OwnerId)
		if err != nil {
			return nil, fmt.Errorf("failed to get store : %w", err)
		}
		stores = append(stores, store)
	}

	return stores, nil
}

func (repo *storeRepository) GetStoreByName(name string) (*entity.Store, error) {
	var store entity.Store
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

func (repo *storeRepository) UpdateStore(updateStore *entity.Store, tx *sql.Tx) (*entity.Store, error) {
	stmt, err := repo.db.Prepare("UPDATE tbl_store SET name = $1, no_hp=$2,email=$3,address=$4 WHERE id = $5")
	if err != nil {
		return nil, fmt.Errorf("failed to update store : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateStore.Name, updateStore.NoHp, updateStore.Email, updateStore.Address, updateStore.Id)

	validate(err, "update store", tx)

	return updateStore, nil
}

func (repo *storeRepository) UpdateMainStore(newUpdateStore *webstore.StoreUpdateRequest) (*webstore.StoreCreateResponse, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		panic(err)
	}
	newUpdateStoreNoHp, err := strconv.Atoi(newUpdateStore.NoHp)
	if err != nil {
		return nil, err
	}
	updateStore := entity.Store{
		Id:      newUpdateStore.Id,
		Name:    newUpdateStore.Name,
		NoHp:    newUpdateStoreNoHp,
		Email:   newUpdateStore.Email,
		Address: newUpdateStore.Address,
		OwnerId: newUpdateStore.OwnerId,
	}
	updateBank := entity.Bank{
		Id:          newUpdateStore.BankId,
		Name:        newUpdateStore.BankName,
		BankAccount: newUpdateStore.BankAccount,
		AccountName: newUpdateStore.AccountName,
	}

	store, err := repo.UpdateStore(&updateStore, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to update store : %w", err)
	}

	bank, err := repo.bankStoreRepository.UpdateBankStore(&updateBank, tx)
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

func (repo *storeRepository) DeleteStore(id int, tx *sql.Tx) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_store SET is_deleted = true WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to delete store : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	validate(err, "delete store", tx)

	return nil
}

func (repo *storeRepository) DeleteMainStore(storeid int, ownerId int) error {
	tx, err := repo.db.Begin()
	if err != nil {
		panic(err)
	}

	err = repo.DeleteStore(storeid, tx)
	if err != nil {
		return fmt.Errorf("failed to delete store : %w", err)
	}

	getBankStoreByStoreId, err := repo.bankStoreRepository.GetBankStoreByStoreId(storeid)
	if err != nil {
		return fmt.Errorf("failed to get bank store : %w", err)
	}
	for _, v := range getBankStoreByStoreId {
		err = repo.bankStoreRepository.DeleteBank(v.Id, tx)
		if err != nil {
			return fmt.Errorf("failed to delete bank : %w", err)
		}
	}

	err = repo.bankStoreRepository.DeleteBankStore(storeid, tx)
	if err != nil {
		return fmt.Errorf("failed to delete bank store : %w", err)
	}

	err = repo.productRepository.DeleteProductByStoreId(storeid, tx)
	if err != nil {
		return fmt.Errorf("failed to delete product : %w", err)
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		return fmt.Errorf("failed to delete store : %w", errCommit)
	}

	return nil
}
func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "transaction rollback")
	} else {
		fmt.Println("success")
	}
}
