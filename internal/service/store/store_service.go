package storeservice

import (
	"pancakaki/internal/domain/entity"
	webstore "pancakaki/internal/domain/web/store"
	storerepository "pancakaki/internal/repository/store"
)

type StoreService interface {
	// CreateStore(newStore *entity.Store, tx *sql.Tx) (*entity.Store, error)
	GetStoreByOwnerId(id int) (*entity.Store, error)
	GetStoreByName(name string) (*entity.Store, error)
	CreateMainStore(newTransactionStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error)
	UpdateMainStore(newUpdateStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error)
}

type storeService struct {
	storeRepo storerepository.StoreRepository
}

func (s *storeService) CreateMainStore(newTransactionStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error) {
	return s.storeRepo.CreateMainStore(newTransactionStore)
}

func (s *storeService) GetStoreByOwnerId(id int) (*entity.Store, error) {
	return s.storeRepo.GetStoreByOwnerId(id)
}

func (s *storeService) GetStoreByName(name string) (*entity.Store, error) {
	return s.storeRepo.GetStoreByName(name)
}

func (s *storeService) UpdateMainStore(newUpdateStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error) {
	return s.storeRepo.UpdateMainStore(newUpdateStore)
}
func NewStoreService(storeRepo storerepository.StoreRepository) StoreService {
	return &storeService{storeRepo: storeRepo}
}
