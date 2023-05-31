package storeservice

import (
	"errors"
	"pancakaki/internal/domain/entity"
	webstore "pancakaki/internal/domain/web/store"
	bankstorerepository "pancakaki/internal/repository/bank_store"
	storerepository "pancakaki/internal/repository/store"
	transactionrepository "pancakaki/internal/repository/transaction"
	"strconv"
)

type StoreService interface {
	// CreateStore(newStore *entity.Store, tx *sql.Tx) (*entity.Store, error)
	GetStoreByOwnerId(id int) ([]entity.Store, error)
	GetStoreByName(name string) (*entity.Store, error)
	GetTransactionByStoreIdAndOwnerId(storeId int, ownerId int) ([]entity.TransactionStore, error)
	CreateMainStore(newTransactionStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error)
	UpdateMainStore(newUpdateStore *webstore.StoreUpdateRequest) (*webstore.StoreCreateResponse, error)
	UpdatePayment(newUpdateTransaction *entity.TransactionOrderDetail, storeId int, ownerId int) (*entity.TransactionOrderDetail, error)
	DeleteMainStore(storeid int, ownerId int) error
}

type storeService struct {
	storeRepo       storerepository.StoreRepository
	bankRepo        bankstorerepository.BankStoreRepository
	transactionRepo transactionrepository.TransactionRepository
}

func (s *storeService) CreateMainStore(newTransactionStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error) {

	return s.storeRepo.CreateMainStore(newTransactionStore)
}

func (s *storeService) GetStoreByOwnerId(id int) ([]entity.Store, error) {
	return s.storeRepo.GetStoreByOwnerId(id)
}

func (s *storeService) GetStoreByName(name string) (*entity.Store, error) {
	return s.storeRepo.GetStoreByName(name)
}

func (s *storeService) GetTransactionByStoreIdAndOwnerId(storeId int, ownerId int) ([]entity.TransactionStore, error) {
	getStoreByOwnerId, err := s.storeRepo.GetStoreByOwnerId(ownerId)
	storeIdStr := strconv.Itoa(storeId)
	if err != nil {
		return nil, errors.New("store with id " + storeIdStr + " not found")
	}

	checkStoreId := false
	for _, v := range getStoreByOwnerId {
		if v.Id == storeId {
			checkStoreId = true
		}
	}
	if !checkStoreId {
		return nil, errors.New("store with id " + storeIdStr + " is unauthorized")
	}
	return s.storeRepo.GetTransactionByStoreIdAndOwnerId(storeId, ownerId)
}

func (s *storeService) UpdateMainStore(newUpdateStore *webstore.StoreUpdateRequest) (*webstore.StoreCreateResponse, error) {

	getStoreByOwnerId, err := s.storeRepo.GetStoreByOwnerId(newUpdateStore.OwnerId)
	newUpdateStoreIdStr := strconv.Itoa(newUpdateStore.Id)
	if err != nil {
		return nil, errors.New("store with id " + newUpdateStoreIdStr + " not found")
	}

	checkStoreId := false
	for _, v := range getStoreByOwnerId {
		if v.Id == newUpdateStore.Id {
			checkStoreId = true
		}
	}
	if !checkStoreId {
		return nil, errors.New("store with id " + newUpdateStoreIdStr + " is unauthorized")
	}

	// fmt.Println(newUpdateStore.Id)
	getBankStoreByStoreId, err := s.bankRepo.GetBankStoreByStoreId(newUpdateStore.Id)
	newUpdateBankId := strconv.Itoa(newUpdateStore.BankId)
	if err != nil {
		return nil, errors.New("bank with id " + newUpdateBankId + " not found")
	}
	// fmt.Println(getBankStoreByStoreId)
	checkBankStoreId := false
	for _, v := range getBankStoreByStoreId {
		if v.Id == newUpdateStore.BankId {
			checkBankStoreId = true
		}
	}
	if !checkBankStoreId {
		return nil, errors.New("bank with id " + newUpdateBankId + " is unauthorized")
	}

	return s.storeRepo.UpdateMainStore(newUpdateStore)
}

func (s *storeService) UpdatePayment(newUpdateTransaction *entity.TransactionOrderDetail, storeId int, ownerId int) (*entity.TransactionOrderDetail, error) {
	getStoreByOwnerId, err := s.storeRepo.GetStoreByOwnerId(ownerId)
	storeIdStr := strconv.Itoa(storeId)
	if err != nil {
		return nil, errors.New("store with id " + storeIdStr + " not found")
	}

	checkStoreId := false
	for _, v := range getStoreByOwnerId {
		if v.Id == storeId {
			checkStoreId = true
		}
	}
	if !checkStoreId {
		return nil, errors.New("store with id " + storeIdStr + " is unauthorized")
	}

	transationStore, err := s.storeRepo.GetTransactionByStoreIdAndOwnerId(storeId, ownerId)
	if err != nil {
		return nil, errors.New("transaction id not found")
	}
	checkTransactionId := false
	for _, v := range transationStore {
		if v.Id == newUpdateTransaction.Id {
			checkTransactionId = true
		}
	}
	transactionIdStr := strconv.Itoa(newUpdateTransaction.Id)
	if !checkTransactionId {
		return nil, errors.New("transaction with id " + transactionIdStr + " is unauthorized")
	}

	transactionOrderDetail, err := s.transactionRepo.FindById(newUpdateTransaction.Id)
	if err != nil {
		return nil, errors.New("transaction not found")
	}
	newUpdateTransaction.BuyDate = transactionOrderDetail.BuyDate
	newUpdateTransaction.Status = "on delivery, transaction completed"
	newUpdateTransaction.TotalPrice = transactionOrderDetail.TotalPrice
	newUpdateTransaction.Photo = transactionOrderDetail.Photo
	newUpdateTransaction.Tax = transactionOrderDetail.Tax
	newUpdateTransaction.VirtualAccount = transactionOrderDetail.VirtualAccount

	return s.storeRepo.UpdatePayment(newUpdateTransaction, storeId, ownerId)
}
func (s *storeService) DeleteMainStore(storeid int, ownerId int) error {
	getStoreByOwnerId, err := s.storeRepo.GetStoreByOwnerId(ownerId)
	storeIdStr := strconv.Itoa(storeid)
	if err != nil {
		return errors.New("store with id " + storeIdStr + " not found")
	}

	checkStoreId := false
	for _, v := range getStoreByOwnerId {
		if v.Id == storeid {
			checkStoreId = true
		}
	}
	if !checkStoreId {
		return errors.New("store with id " + storeIdStr + " is unauthorized")
	}

	return s.storeRepo.DeleteMainStore(storeid, ownerId)
}
func NewStoreService(
	storeRepo storerepository.StoreRepository,
	bankRepo bankstorerepository.BankStoreRepository,
	transactionRepo transactionrepository.TransactionRepository) StoreService {
	return &storeService{
		storeRepo:       storeRepo,
		bankRepo:        bankRepo,
		transactionRepo: transactionRepo}
}
