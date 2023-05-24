package bankservice

import (
	"pancakaki/internal/domain/entity"
	bankstorerepository "pancakaki/internal/repository/bank"
)

type BankService interface {
	GetBankAdminById(id int) ([]entity.Bank, error)
}

type bankService struct {
	bankRepo bankstorerepository.BankStoreRepository
}

func (s *bankService) GetBankAdminById(id int) ([]entity.Bank, error) {
	return s.bankRepo.GetBankAdminById(id)
}

func NewBankService(bankRepo bankstorerepository.BankStoreRepository) BankService {
	return &bankService{bankRepo: bankRepo}
}
