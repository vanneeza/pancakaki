package bankservice

import (
	entity "pancakaki/internal/domain/entity/bank"
	bankrepository "pancakaki/internal/repository/bank"
)

type BankService interface {
	GetBankAdminById(id int) ([]entity.Bank, error)
}

type bankService struct {
	bankRepo bankrepository.BankRepository
}

func (s *bankService) GetBankAdminById(id int) ([]entity.Bank, error) {
	return s.bankRepo.GetBankAdminById(id)
}

func NewBankService(bankRepo bankrepository.BankRepository) BankService {
	return &bankService{bankRepo: bankRepo}
}
