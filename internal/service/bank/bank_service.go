package bankservice

import (
	"pancakaki/internal/domain/entity"
	bankstorerepository "pancakaki/internal/repository/bank_store"
)

type BankService interface {
	GetBankAdminById(id int) ([]entity.Bank, error)
}

type bankService struct {
	bankstorerepository bankstorerepository.BankStoreRepository
}

func (s *bankService) GetBankAdminById(id int) ([]entity.Bank, error) {
	return s.bankstorerepository.GetBankAdminById(id)
}

func NewBankService(bankstorerepository bankstorerepository.BankStoreRepository) BankService {
	return &bankService{bankstorerepository: bankstorerepository}
}
