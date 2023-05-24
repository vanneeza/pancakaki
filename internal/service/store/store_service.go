package storeservice

import (
	entitystore "pancakaki/internal/domain/entity/store"
	webstore "pancakaki/internal/domain/web/store"
	storerepository "pancakaki/internal/repository/store"
)

type StoreService interface {
	// CreateStore(newStore *entitystore.Store, tx *sql.Tx) (*entitystore.Store, error)
	GetStoreByOwnerId(id int) (*entitystore.Store, error)
	GetStoreByName(name string) (*entitystore.Store, error)
	CreateMainStore(newTransactionStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error)
	UpdateMainStore(newUpdateStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error)
}

type storeService struct {
	storeRepo storerepository.StoreRepository
}

func (s *storeService) CreateMainStore(newTransactionStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error) {
	return s.storeRepo.CreateMainStore(newTransactionStore)
}

func (s *storeService) GetStoreByOwnerId(id int) (*entitystore.Store, error) {
	return s.storeRepo.GetStoreByOwnerId(id)
}

func (s *storeService) GetStoreByName(name string) (*entitystore.Store, error) {
	return s.storeRepo.GetStoreByName(name)
}

func (s *storeService) UpdateMainStore(newUpdateStore *webstore.StoreCreateRequest) (*webstore.StoreCreateResponse, error) {
	return s.storeRepo.UpdateMainStore(newUpdateStore)
}
func NewStoreService(storeRepo storerepository.StoreRepository) StoreService {
	return &storeService{storeRepo: storeRepo}
}
